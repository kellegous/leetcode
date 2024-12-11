package p2779

func maximumBeauty(nums []int, k int) int {
	n := len(nums)
	minNum := nums[0]
	maxNum := nums[0]
	for i := 1; i < n; i++ {
		minNum = min(nums[i], minNum)
		maxNum = max(nums[i], maxNum)
	}

	maxTotal := maxNum + k
	minTotal := minNum - k
	freq := map[int]int{}
	for _, num := range nums {
		freq[num-k]++
		freq[num+k+1]--
	}

	var cb, mb int
	for i := minTotal; i <= maxTotal; i++ {
		cb += freq[i]
		mb = max(mb, cb)
	}
	return mb
}
