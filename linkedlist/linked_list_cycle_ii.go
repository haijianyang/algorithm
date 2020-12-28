package linkedlist

/*
https://leetcode-cn.com/problems/linked-list-cycle-ii
*/

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	fast := head.Next
	slow := head

	for fast != nil && fast.Next != nil && slow != nil {
		if fast == slow {

		}

		fast = head.Next.Next
		slow = slow.Next
	}

	return nil
}
