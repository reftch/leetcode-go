package ds

// ListNode is the shared singly-linked list node used by LeetCode problems.
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode builds a linked list from vals and returns its head.
// Returns nil for an empty slice.
func NewListNode(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// ToSlice converts a linked list back to a slice. Returns nil for nil head.
func (n *ListNode) ToSlice() []int {
	var out []int
	for cur := n; cur != nil; cur = cur.Next {
		out = append(out, cur.Val)
	}
	return out
}

// Equal reports whether two lists contain the same values in order.
func (n *ListNode) Equal(other *ListNode) bool {
	a, b := n, other
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}
