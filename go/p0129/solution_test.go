package p0129

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input    *TreeNode
		Expected int
	}{
		{newTree(1, 2, 3), 25},
		{newTree(4, 9, 0, 5, 1), 1026},
	}

	for _, test := range tests {
		got := sumNumbers(test.Input)
		if got != test.Expected {
			t.Fatalf(
				"Input: %s, Expected: %d, Got: %d",
				test.Input.String(),
				test.Expected,
				got)
		}
	}
}
