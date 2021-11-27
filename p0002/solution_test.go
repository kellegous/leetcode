package p0002

import (
	"strconv"
	"strings"
	"testing"
)

func listAreSame(a, b *ListNode) bool {
	if a == nil {
		return b == nil
	} else if b == nil {
		return false
	}
	return a.Val == b.Val && listAreSame(a.Next, b.Next)
}

func listToString(n *ListNode) string {
	var vals []string
	for ; n != nil; n = n.Next {
		vals = append(vals, strconv.Itoa(n.Val))
	}
	return strings.Join(vals, ", ")
}

func newList(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	return &ListNode{
		Val:  vals[0],
		Next: newList(vals[1:]...),
	}
}

func TestSolution(t *testing.T) {
	tests := []struct {
		L1       *ListNode
		L2       *ListNode
		Expected *ListNode
	}{
		{newList(2, 4, 3), newList(5, 6, 4), newList(7, 0, 8)},
		{newList(0), newList(0), newList(0)},
		{newList(9, 9, 9, 9, 9, 9, 9), newList(9, 9, 9, 9), newList(8, 9, 9, 9, 0, 0, 0, 1)},
	}

	for _, test := range tests {
		got := addTwoNumbers(test.L1, test.L2)
		if !listAreSame(got, test.Expected) {
			t.Fatalf(
				"L1: [%s], L2: [%s], Got: [%s], Expected: [%s]",
				listToString(test.L1),
				listToString(test.L2),
				listToString(got),
				listToString(test.Expected))
		}
	}
}
