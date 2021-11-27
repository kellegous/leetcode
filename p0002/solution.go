package p0002

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return add(l1, l2, 0)
}

func add(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry == 0 {
			return nil
		}
		return &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	var v1, v2 int
	var n1, n2 *ListNode

	if l1 != nil {
		v1 = l1.Val
		n1 = l1.Next
	}

	if l2 != nil {
		v2 = l2.Val
		n2 = l2.Next
	}

	v := v1 + v2 + carry

	if v >= 10 {
		return &ListNode{
			Val:  v - 10,
			Next: add(n1, n2, 1),
		}
	}

	return &ListNode{
		Val:  v,
		Next: add(n1, n2, 0),
	}
}
