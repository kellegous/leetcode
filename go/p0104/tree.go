package p0104

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func valueOf(v interface{}) *int {
	switch t := v.(type) {
	case int:
		return &t
	default:
		return nil
	}
}

func newTree(vals ...interface{}) *TreeNode {
	return newNode(vals, 0)
}

func newNode(vals []interface{}, ix int) *TreeNode {
	if ix >= len(vals) {
		return nil
	}

	v := valueOf(vals[ix])
	if v == nil {
		return nil
	}

	return &TreeNode{
		Val:   *v,
		Left:  newNode(vals, 2*ix+1),
		Right: newNode(vals, 2*ix+2),
	}
}

func (n *TreeNode) String() string {
	if n == nil {
		return "nil"
	}
	if n.Left == nil && n.Right == nil {
		return fmt.Sprintf("(%d)", n.Val)
	}

	return fmt.Sprintf(
		"(%d %s %s)",
		n.Val, n.Left.String(),
		n.Right.String())
}
