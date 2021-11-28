package p0129

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	return sum(root, 0)
}

func sum(root *TreeNode, acc int) int {
	if root == nil {
		return 0
	}

	r, l, v := root.Left, root.Right, root.Val

	if r == nil && l == nil {
		return acc + root.Val
	}

	res := (acc + v) * 10
	var s int

	if l != nil {
		s += sum(l, res)
	}

	if r != nil {
		s += sum(r, res)
	}

	return s
}
