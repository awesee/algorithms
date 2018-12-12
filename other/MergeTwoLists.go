package other

// 21. Merge Two Sorted Lists
//
// Favorite
// Share
// Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
//
// Example:
//
// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	l := &ListNode{}
	t := l
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			t.Next = l1
			l1 = l1.Next
		} else {
			t.Next = l2
			l2 = l2.Next
		}
		t = t.Next
	}
	if l1 == nil {
		t.Next = l2
	} else {
		t.Next = l1
	}

	return l.Next
}
