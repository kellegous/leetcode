package p2779

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		Nums     []int
		K        int
		Expected int
	}{
		{[]int{4, 6, 1, 2}, 2, 3},
		{[]int{1, 1, 1, 1}, 10, 4},
		{[]int{49, 26}, 12, 2},
		{[]int{100000}, 0, 1},
	}

	for _, test := range tests {
		if got := maximumBeauty(test.Nums, test.K); got != test.Expected {
			t.Fatalf("test failed with nums = %#v and k = %#v, expected %d, got %d",
				test.Nums,
				test.K,
				test.Expected,
				got)
		}
	}
}
