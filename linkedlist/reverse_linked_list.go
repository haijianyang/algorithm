package linkedlist

/*
https://leetcode-cn.com/problems/reverse-linked-list
*/
func reverseList(head *ListNode) *ListNode {
	var start *ListNode
	for head != nil {
		temp := head.Next
		head.Next = start
		start = head
		head = temp
	}

	return start
}
