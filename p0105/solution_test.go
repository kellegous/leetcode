package p0105

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Preorder []int
		Inorder  []int
		Expected *TreeNode
	}{
		{
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
			newTree(3, 9, 20, nil, nil, 15, 7),
		},
		{
			[]int{-1},
			[]int{-1},
			newTree(-1),
		},
	}

	for _, test := range tests {
		got := buildTree(test.Preorder, test.Inorder)
		if !nodesAreSame(got, test.Expected) {
			t.Fatalf(
				"Preorder: %v, Inorder: %v, Expected: %s, Got: %s",
				test.Preorder,
				test.Inorder,
				test.Expected.String(),
				got.String())
		}
	}
}
