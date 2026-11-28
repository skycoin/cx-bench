-- bench.lua runs the cx-bench programs in Lua 5.4 or later. See
-- ../../README.md for the protocol, which every language follows.
--
--	lua bench.lua <name|number> <size> [-time=ms|us]   one run of two passes
--	lua bench.lua empty                                 start and exit only
--	lua bench.lua -server                               benchmarks from stdin
--
-- Lua has no monotonic wall clock in its standard library, so the times it
-- prints come from os.clock, the CPU time of the process.
io.stderr:write("S")

local NAMES = { "count", "fib", "array", "struct", "memtest", "primes", "ackermann",
	"digitalroot", "factorialiter", "factorialrec", "structs", "structserialize", "fibiter",
	"ackermanniter" }
local FNV_OFFSET = 0xcbf29ce484222325 -- 14695981039346656037
local FNV_PRIME = 1099511628211

-- fnv mixes one 64 bit word. Lua integers wrap like a uint64.
local function fnv(h, v)
	return (h ~ v) * FNV_PRIME
end

local function unsigned(h)
	if h >= 0 then
		return string.format("%d", h)
	end
	local q = (h >> 1) // 5
	return string.format("%d%d", q, h - q * 10)
end

local dir = arg[0]:match("^(.*)/") or "."
local function load(number)
	return dofile(dir .. "/" .. NAMES[number] .. ".lua")(fnv, FNV_OFFSET)
end

local function find(key)
	local number = math.tointeger(tonumber(key))
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
	io.stderr:write("usage: lua bench.lua <name|number> <size> [-time=ms|us] | empty | -server\n")
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
			local start = os.clock()
			local h, ok = runs[number](math.tointeger(tonumber(size)))
			local spent = elapsed(os.clock() - start, unit)
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
	local size = math.tointeger(tonumber(arg[2]))
	local unit = arg[3] and arg[3]:match("^%-time=(%a+)$")
	if not number or not size or size < 0 or (arg[3] and unit ~= "ms" and unit ~= "us") then
		usage()
	end
	local run = load(number)
	local start = os.clock()
	local h1, ok1 = run(size)
	io.stderr:write("1")
	local middle = os.clock()
	local h2, ok2 = run(size)
	io.stderr:write("2")
	local stop = os.clock()
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

os.exit(main())
