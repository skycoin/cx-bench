-- Ackermann(3, 5) with an explicit stack, 42,438 steps, once per round. The
-- loop variant of ackermann. Each step is one call of the recursive form: pop
-- m, then push what the call would recurse into.
-- zero is the 0 the compiled forms read so no compiler precomputes a round, see
-- README.md. It costs the same read here.
local zero = 0

return function(fnv, offset)
	local stack = {}
	local function ackermann(m, n)
		local top = 1
		stack[1] = m
		while top > 0 do
			m = stack[top]
			top = top - 1
			if m == 0 then
				n = n + 1
			elseif n == 0 then
				top = top + 1
				stack[top] = m - 1
				n = 1
			else
				stack[top + 1] = m - 1
				stack[top + 2] = m
				top = top + 2
				n = n - 1
			end
		end
		return n
	end
	return function(size)
		local h, ok = offset, true
		for _ = 1, size do
			local r = ackermann(3, 5 + zero)
			h = fnv(h, r)
			ok = ok and r == 253
		end
		return h, ok
	end
end
