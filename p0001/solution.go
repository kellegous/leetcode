package p0001

func twoSum(nums []int, target int) []int {
	idx := map[int]int{}
	for i, n := range nums {
		if j, ok := idx[target-n]; ok {
			return []int{j, i}
		}
		idx[n] = i
	}
	return nil // ??
}
