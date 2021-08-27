package list_node

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 反转链表: https://leetcode-cn.com/problems/reverse-linked-list/
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

/**
 * K 个一组反转链表：https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	return nil
}
