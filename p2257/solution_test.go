package p2257

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		M        int
		N        int
		Guards   [][]int
		Walls    [][]int
		Expected int
	}{
		{
			4, 6,
			[][]int{
				{0, 0},
				{1, 1},
				{2, 3},
			},
			[][]int{
				{0, 1},
				{2, 2},
				{1, 4},
			},
			7,
		},
		{
			3, 3,
			[][]int{
				{1, 1},
			},
			[][]int{
				{0, 1},
				{1, 0},
				{2, 1},
				{1, 2},
			},
			4,
		},
	}

	for _, test := range tests {
		if got := countUnguarded(
			test.M,
			test.N,
			test.Guards,
			test.Walls,
		); got != test.Expected {
			t.Fatalf("For test m = %d, n = %d, guards = %v, walls = %v: expected %d got %d",
				test.M,
				test.N,
				test.Guards,
				test.Walls,
				test.Expected,
				got)
		}
	}
}
