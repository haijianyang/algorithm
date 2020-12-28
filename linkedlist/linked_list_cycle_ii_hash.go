package linkedlist

/*
https://leetcode-cn.com/problems/linked-list-cycle-ii
*/

func detectCycle(head *ListNode) *ListNode {
	node := head
	hash := make(map[*ListNode]*ListNode)
	for node != nil {
		if _, ok := hash[node]; ok {
			return node
		}

		hash[node] = node
		node = node.Next
	}

	return nil
}
