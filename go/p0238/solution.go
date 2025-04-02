package p0238

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	n := len(nums)
	for i, acc := 0, 1; i < n; i++ {
		res[i] = acc
		acc *= nums[i]
	}
	for i, acc := n-1, 1; i >= 0; i-- {
		res[i] *= acc
		acc *= nums[i]
	}
	return res
}
