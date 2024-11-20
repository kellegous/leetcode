package p2461

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		Nums     []int
		K        int
		Expected int64
	}{
		{[]int{1, 5, 4, 2, 9, 9, 9}, 3, 15},
		{[]int{4, 4, 4}, 3, 0},
		{[]int{1, 2, 2, 3, 4}, 3, 9},
	}

	for _, test := range tests {
		if got := maximumSubarraySum(test.Nums, test.K); got != test.Expected {
			t.Fatalf(
				"For test with nums = %v, k = %v: expected %d got %d",
				test.Nums,
				test.K,
				test.Expected,
				got,
			)
		}
	}
}
