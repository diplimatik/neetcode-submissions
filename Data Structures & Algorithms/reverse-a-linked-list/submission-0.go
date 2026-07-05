/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		if tmp == nil {
			break
		}
		head = tmp
	}

	return head
}
