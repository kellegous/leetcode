package p0106

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		InOrder   []int
		PostOrder []int
		Expected  *TreeNode
	}{
		{
			[]int{9, 3, 15, 20, 7},
			[]int{9, 15, 7, 20, 3},
			newTree(3, 9, 20, nil, nil, 15, 7),
		},
		{
			[]int{-1},
			[]int{-1},
			newTree(-1),
		},
		{
			[]int{9, 15, 20, 7, 3},
			[]int{9, 15, 7, 20, 3},
			newTree(3, 20, nil, 15, 7, nil, nil, 9),
		},
	}

	for _, test := range tests {
		got := buildTree(test.InOrder, test.PostOrder)
		if !nodesAreSame(got, test.Expected) {
			t.Fatalf(
				"Inorder: %v, Postorder: %v, Expected: %s, Got: %s",
				test.InOrder,
				test.PostOrder,
				test.Expected.String(),
				got.String())
		}
	}
}
