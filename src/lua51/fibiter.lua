-- fib(25) with a loop, 25 steps, once per round. The loop variant of fib.
-- zero is the 0 the compiled forms read so no compiler precomputes a round, see
-- README.md. It costs the same read here.
local zero = 0

return function(fnv, offset)
	local function fib(n)
		local a, b = 0, 1
		-- gopher-lua v1.1.2 gets a, b = b, a + b wrong (it reads the new a), so
		-- the sum goes through a local.
		for _ = 1, n do
			local sum = a + b
			a = b
			b = sum
		end
		return a
	end
	return function(size)
		local h, ok = offset(), true
		for _ = 1, size do
			local r = fib(25 + zero)
			h = fnv(h, r)
			ok = ok and r == 75025
		end
		return h, ok
	end
end
