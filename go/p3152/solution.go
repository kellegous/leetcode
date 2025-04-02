package p3152

func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
	// edges is an array that contains the number of adjacent pairs that have
	// matching parity up until that point.
	edges := make([]int, n)
	for i := 1; i < n; i++ {
		if nums[i-1]&1 == nums[i]&1 {
			edges[i] = 1
		}
		edges[i] += edges[i-1]
	}

	res := make([]bool, len(queries))
	for i, query := range queries {
		res[i] = (edges[query[1]] - edges[query[0]]) == 0
	}

	return res
}
