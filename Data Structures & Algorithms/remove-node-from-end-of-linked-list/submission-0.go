	/**
	* Definition for singly-linked list.
	* type ListNode struct {
	*     Val int
	*     Next *ListNode
	* }
	*/


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	i := 0
	for currHead := head; currHead != nil; currHead = currHead.Next {
		i++
	}
	i -= n

	if i == 0 {
		return head.Next
	}

	for currHead := head; i > 0; currHead = currHead.Next {
		if i == 1 {
			currHead.Next = currHead.Next.Next
		}
		i--
	}
	return head
}

