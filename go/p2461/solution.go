package p2461

func maximumSubarraySum(nums []int, k int) int64 {
	var m int64
	seen := map[int]int{}
	var sum int64
	for i, n := 0, len(nums); i < n; i++ {
		// remove the value falling out of the k window
		if ix := i - k; ix >= 0 {
			num := nums[i-k]
			seen[num]--
			if seen[num] == 0 {
				delete(seen, num)
			}
			sum -= int64(num)
		}

		// add the value coming into the k window
		seen[nums[i]]++
		sum += int64(nums[i])

		// if k values are in the set, the sum can be considered
		if len(seen) == k {
			m = max(m, sum)
		}
	}
	return m
}
