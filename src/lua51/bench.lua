-- bench.lua runs the cx-bench programs in Lua 5.1, for gopher-lua. See
-- ../../README.md for the protocol, which every language follows.
--
--	gopher-lua bench.lua <name|number> <size> [-time=ms|us]
--	gopher-lua bench.lua empty
--	gopher-lua bench.lua -server
--
-- Lua 5.1 numbers are doubles with no bit operators, so the hash is eight 8
-- bit limbs, low first, and XOR is a byte table lookup. primes,
-- factorialiter and factorialrec need 64 bit integers and are left out. The
-- times come from os.clock.
io.stderr:write("S")

local NAMES = { "count", "fib", "array", "struct", "memtest", "primes", "ackermann",
	"digitalroot", "factorialiter", "factorialrec", "structs", "structserialize", "fibiter",
	"ackermanniter" }
local MISSING = { primes = true, factorialiter = true, factorialrec = true }

local XOR = {}
for b = 0, 255 do
	XOR[b + 1] = b
end
for a = 1, 255 do
	local half = (a - a % 2) / 2 * 256 + 1
	for b = 0, 255 do
		XOR[a * 256 + b + 1] = 2 * XOR[half + (b - b % 2) / 2] + (a + b) % 2
	end
end

-- offset returns a new hash at the offset basis 14695981039346656037.
local function offset()
	return { 0x25, 0x23, 0x22, 0x84, 0xe4, 0x9c, 0xf2, 0xcb }
end

-- fnv mixes one word from 0 to 2^53 into the hash h: XOR, then times the
-- prime 2^40 + 0x1b3 mod 2^64.
local function fnv(h, v)
	for i = 1, 8 do
		local b = v % 256
		v = (v - b) / 256
		h[i] = XOR[h[i] * 256 + b + 1]
	end
	for k = 8, 1, -1 do
		local t = h[k] * 0xb3
		if k >= 2 then
			t = t + h[k - 1]
		end
		if k >= 6 then
			t = t + h[k - 5]
		end
		h[k] = t
	end
	local carry = 0
	for k = 1, 8 do
		local t = h[k] + carry
		local b = t % 256
		h[k] = b
		carry = (t - b) / 256
	end
	return h
end

local function decimal(h)
	local limbs = { h[1], h[2], h[3], h[4], h[5], h[6], h[7], h[8] }
	local digits = ""
	repeat
		local rest, nonzero = 0, false
		for k = 8, 1, -1 do
			local t = rest * 256 + limbs[k]
			limbs[k] = (t - t % 10) / 10
			rest = t % 10
			nonzero = nonzero or limbs[k] ~= 0
		end
		digits = rest .. digits
	until not nonzero
	return digits
end

local dir = arg[0]:match("^(.*)/") or "."
local function load(number)
	return dofile(dir .. "/" .. NAMES[number] .. ".lua")(fnv, offset)
end

local function integer(text)
	local n = tonumber(text)
	if n and n == math.floor(n) then
		return n
	end
end

local function find(key)
	local number = integer(key)
	if not number then
		for i, name in ipairs(NAMES) do
			if name == key then
				number = i
			end
		end
	end
	if number and number >= 1 and number <= #NAMES and not MISSING[NAMES[number]] then
		return number
	end
end

local function elapsed(seconds, unit)
	return string.format("%.0f", math.floor(seconds * (unit == "ms" and 1e3 or 1e6)))
end

local function usage()
	io.stderr:write("usage: gopher-lua bench.lua <name|number> <size> [-time=ms|us] | empty | -server\n")
	os.exit(2)
end

local function server()
	local runs = {}
	while true do
		local line = io.read("*l")
		if not line then
			break
		end
		local number, size, unit = line:match("^%s*(%d+)%s*(%d*)%s*(%a*)")
		number = tonumber(number)
		if number == 0 then
			break
		end
		if number and find(number) and size ~= "" and unit ~= "" then
			runs[number] = runs[number] or load(number)
			local start = os.clock()
			local h, ok = runs[number](tonumber(size))
			local spent = elapsed(os.clock() - start, unit)
			io.write(number, " ", decimal(h), " ", spent, ok and "" or " FAIL", "\n")
		else
			io.write(tostring(number), " error\n")
		end
		io.stdout:flush()
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
	local start = os.clock()
	local h1, ok1 = run(size)
	io.stderr:write("1")
	local middle = os.clock()
	local h2, ok2 = run(size)
	io.stderr:write("2")
	local stop = os.clock()
	local d1, d2 = decimal(h1), decimal(h2)
	local ok = ok1 and ok2 and d1 == d2
	print(d1)
	print(NAMES[number] .. (ok and " pass" or " FAIL"))
	if unit then
		print("time " .. unit .. " loop1=" .. elapsed(middle - start, unit) .. " loop2=" .. elapsed(stop - middle, unit))
	end
	io.stdout:flush()
	io.stderr:write("F")
	return ok and 0 or 1
end

os.exit(main())
