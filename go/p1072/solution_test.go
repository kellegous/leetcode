package p1072

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		Matrix   [][]int
		Expected int
	}{
		{
			[][]int{{0, 1}, {1, 1}},
			1,
		},
		{
			[][]int{{0, 1}, {1, 0}},
			2,
		},
		{
			[][]int{
				{0, 0, 0},
				{0, 0, 1},
				{1, 1, 0},
			},
			2,
		},
		{
			[][]int{
				{1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1},
				{1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0},
				{1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1},
				{1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0},
				{1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1},
			},
			2,
		},
	}

	for _, test := range tests {
		if got := maxEqualRowsAfterFlips(test.Matrix); got != test.Expected {
			t.Fatalf(
				"For text w/ matrix = %v got %d but expected %d",
				test.Matrix,
				got,
				test.Expected,
			)
		}
	}
}
