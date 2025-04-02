package p0404

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input    *TreeNode
		Expected int
	}{
		{newTree(3, 9, 20, nil, nil, 15, 7), 24},
		{newTree(1), 0},
	}

	for _, test := range tests {
		got := sumOfLeftLeaves(test.Input)
		if got != test.Expected {
			t.Fatalf(
				"Input: %s, Expected: %d, Got: %d",
				test.Input.String(),
				test.Expected,
				got)
		}
	}
}
