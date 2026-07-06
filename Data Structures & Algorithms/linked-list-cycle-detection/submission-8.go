/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	fastHead := head
	for head != nil && fastHead.Next != nil && fastHead.Next.Next != nil {
		fastHead = fastHead.Next.Next
		if head == fastHead {
			return true
		}
		head = head.Next
	}
	return false
}
