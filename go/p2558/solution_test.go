package p2558

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		Gifts    []int
		K        int
		Expected int64
	}{
		{
			[]int{25, 64, 9, 4, 100},
			4,
			29,
		},
		{
			[]int{1, 1, 1, 1},
			4,
			4,
		},
	}

	for _, test := range tests {
		if got := pickGifts(test.Gifts, test.K); got != test.Expected {
			t.Fatalf(
				"test failed w/ gifts = %#v, k = %#v: expected %d but got %d",
				test.Gifts,
				test.K,
				test.Expected,
				got)
		}
	}
}
