package linkedlist

/*
https://leetcode-cn.com/problems/merge-two-sorted-lists
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{}
	cursor := list
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cursor.Next = l1
			cursor = l1
			l1 = l1.Next
		} else {
			cursor.Next = l2
			cursor = l2
			l2 = l2.Next
		}
	}

	for l1 != nil {
		cursor.Next = l1
		cursor = l1
		l1 = l1.Next
	}

	for l2 != nil {
		cursor.Next = l2
		cursor = l2
		l2 = l2.Next
	}

	return list.Next
}
