package p0104

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		Input    *TreeNode
		Expected int
	}{
		{newTree(3, 9, 20, nil, nil, 15, 7), 3},
		{newTree(1, nil, 2), 2},
		{nil, 0},
		{newTree(0), 1},
	}

	for _, test := range tests {
		got := maxDepth(test.Input)
		if got != test.Expected {
			t.Fatalf(
				"Input: %s, Expected: %d, Got: %d",
				test.Input.String(),
				test.Expected,
				got)
		}
	}
}
