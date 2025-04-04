return {
	two_sum = function(nums, target)
		local num_map = {}
		for i, num in ipairs(nums) do
			local j = target - num
			if num_map[target - num] then
				return {num_map[j], i}
			end
			num_map[num] = i
		end
		return nil -- ??
	end,
}