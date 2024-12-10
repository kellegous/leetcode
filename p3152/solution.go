package p3152

import (
	"sort"
)

func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
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

func isArraySpecialX(nums []int, queries [][]int) []bool {
	var spans [][]int
	i := 1
	j := 0
	for n := len(nums); i < n; i++ {
		k := i - 1
		if (nums[i] & 1) == (nums[k] & 1) {
			if k != j {
				spans = append(spans, []int{j, k})
			}
			j = i
		}
	}

	if k := len(nums) - 1; j != k {
		spans = append(spans, []int{j, k})
	}

	res := make([]bool, len(queries))
	for i, query := range queries {
		if query[0] == query[1] {
			res[i] = true
			continue
		}
		// find the spans where 0th item is <= query[0]
		ix := sort.Search(len(spans), func(i int) bool {
			return spans[i][1] >= query[1]
		})
		res[i] = ix < len(spans) && spans[ix][0] <= query[0]
	}

	return res
}
