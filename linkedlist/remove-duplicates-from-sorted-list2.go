package linkedlist

/*
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
*/

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	list := &ListNode{}
	pointer := list
	val := head.Val - 1
	cursor := head
	for cursor != nil {
		if cursor.Val == val {
			cursor = cursor.Next
		} else {
			pointer = cursor
			val = cursor.val
			cursor = cursor.Next
		}
	}

	return list.Next
}
