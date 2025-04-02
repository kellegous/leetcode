package p0105

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	idx := map[int]int{}
	for i, val := range inorder {
		idx[val] = i
	}

	_, n := build(preorder, idx, 0, len(inorder)-1)
	return n
}

func build(po []int, idx map[int]int, i, j int) ([]int, *TreeNode) {
	if i > j {
		return po, nil
	}

	v := po[0]
	r := &TreeNode{Val: v}
	po = po[1:]

	if i != j {
		ix := idx[v]
		po, r.Left = build(po, idx, i, ix-1)
		po, r.Right = build(po, idx, ix+1, j)
	}

	return po, r
}
