package p0404

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var sum int
	if l := root.Left; l != nil {
		if l.Left == nil && l.Right == nil {
			sum = l.Val
		} else {
			sum = sumOfLeftLeaves(l)
		}
	}

	return sum + sumOfLeftLeaves(root.Right)
}
