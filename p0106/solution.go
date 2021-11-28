package p0106

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	idx := map[int]int{}
	for i, val := range inorder {
		idx[val] = i
	}

	_, n := build(postorder, idx, 0, len(inorder)-1)
	return n
}

func build(po []int, idx map[int]int, i, j int) ([]int, *TreeNode) {
	if i > j {
		return po, nil
	}

	n := len(po)
	v := po[n-1]
	r := &TreeNode{Val: v}
	po = po[:n-1]

	if i != j {
		ix := idx[v]
		po, r.Right = build(po, idx, ix+1, j)
		po, r.Left = build(po, idx, i, ix-1)
	}

	return po, r
}
