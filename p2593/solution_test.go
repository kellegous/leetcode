package p2593

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		Nums     []int
		Expected int64
	}{
		{[]int{2, 1, 3, 4, 5, 2}, 7},
		{[]int{2, 3, 5, 1, 3, 2}, 5},
	}

	for _, test := range tests {
		if got := findScore(test.Nums); got != test.Expected {
			t.Fatalf(
				"test failed for nums = %#v, expected %d got %d",
				test.Nums,
				test.Expected,
				got)
		}
	}
}
