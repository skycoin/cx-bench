-- bench.lua runs the cx-bench programs in LuaJIT 2.1. See ../../README.md
-- for the protocol, which every language follows.
--
--	luajit bench.lua <name|number> <size> [-time=ms|us]   one run of two passes
--	luajit bench.lua empty                                 start and exit only
--	luajit bench.lua -server                               benchmarks from stdin
--
-- The hash is a uint64_t cdata, which wraps, and the clock is
-- clock_gettime(CLOCK_MONOTONIC) through the FFI.
io.stderr:write("S")

local bit = require("bit")
local ffi = require("ffi")
ffi.cdef([[
typedef struct { long tv_sec; long tv_nsec; } cxbench_timespec;
int clock_gettime(int clock, cxbench_timespec *t);
]])
local CLOCK_MONOTONIC = 1
local timespec = ffi.new("cxbench_timespec")
local function now()
	ffi.C.clock_gettime(CLOCK_MONOTONIC, timespec)
	return tonumber(timespec.tv_sec) + tonumber(timespec.tv_nsec) * 1e-9
end

local NAMES = { "count", "fib", "array", "struct", "memtest", "primes", "ackermann",
	"digitalroot", "factorialiter", "factorialrec", "structs", "structserialize", "fibiter",
	"ackermanniter" }
local FNV_OFFSET = 0xcbf29ce484222325ULL -- 14695981039346656037
local FNV_PRIME = 1099511628211ULL

local function fnv(h, v)
	return bit.bxor(h, v) * FNV_PRIME
end

local function unsigned(h)
	return (tostring(h):gsub("ULL$", ""))
end

local function integer(text)
	local n = tonumber(text)
	if n and n == math.floor(n) then
		return n
	end
end

local dir = arg[0]:match("^(.*)/") or "."
local function load(number)
	return dofile(dir .. "/" .. NAMES[number] .. ".lua")(fnv, FNV_OFFSET)
end

local function find(key)
	local number = integer(key)
	if number then
		return (number >= 1 and number <= #NAMES) and number or nil
	end
	for i, name in ipairs(NAMES) do
		if name == key then
			return i
		end
	end
end

local function elapsed(seconds, unit)
	return math.floor(seconds * (unit == "ms" and 1e3 or 1e6))
end

local function usage()
	io.stderr:write("usage: luajit bench.lua <name|number> <size> [-time=ms|us] | empty | -server\n")
	os.exit(2)
end

local function server()
	local runs = {}
	for line in io.lines() do
		local number, size, unit = line:match("^%s*(%d+)%s*(%d*)%s*(%a*)")
		number = tonumber(number)
		if number == 0 then
			break
		end
		if number and number >= 1 and number <= #NAMES and size ~= "" and unit ~= "" then
			runs[number] = runs[number] or load(number)
			local start = now()
			local h, ok = runs[number](tonumber(size))
			local spent = elapsed(now() - start, unit)
			io.write(string.format("%d %s %d%s\n", number, unsigned(h), spent, ok and "" or " FAIL"))
		else
			io.write(tostring(number), " error\n")
		end
		io.flush()
	end
	io.stderr:write("F")
end

local function main()
	if arg[1] == "empty" and #arg == 1 then
		io.stderr:write("F")
		return 0
	end
	if arg[1] == "-server" and #arg == 1 then
		server()
		return 0
	end
	if #arg < 2 or #arg > 3 then
		usage()
	end
	local number = find(arg[1])
	local size = integer(arg[2])
	local unit = arg[3] and arg[3]:match("^%-time=(%a+)$")
	if not number or not size or size < 0 or (arg[3] and unit ~= "ms" and unit ~= "us") then
		usage()
	end
	local run = load(number)
	local start = now()
	local h1, ok1 = run(size)
	io.stderr:write("1")
	local middle = now()
	local h2, ok2 = run(size)
	io.stderr:write("2")
	local stop = now()
	local ok = ok1 and ok2 and h1 == h2
	print(unsigned(h1))
	print(NAMES[number] .. (ok and " pass" or " FAIL"))
	if unit then
		print(string.format("time %s loop1=%d loop2=%d", unit, elapsed(middle - start, unit), elapsed(stop - middle, unit)))
	end
	io.stdout:flush()
	io.stderr:write("F")
	return ok and 0 or 1
end

-- true closes the state, so the -jp profiler writes its report.
os.exit(main(), true)
