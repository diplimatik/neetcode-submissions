func reverseKGroup(head *ListNode, k int) *ListNode {
	ansHead, lastNode, tailNode := reverseK(head, k)
	for tailNode != nil {
		newHead, newLastNode, newTailNode := reverseK(tailNode, k)
		if newHead == newLastNode {
			lastNode.Next = newHead
			break
		}
		lastNode.Next = newHead
		lastNode = newLastNode
		tailNode = newTailNode
	}

	return ansHead
}

func reverseK(head *ListNode, k int) (*ListNode, *ListNode, *ListNode) {
	if k == 1 {
		return head, head, head.Next
	}

	if head.Next == nil && k > 1 {
		return head, head, head
	}

	newHead, lastNode, tailNode := reverseK(head.Next, k-1)
	if lastNode == tailNode {
		return head, head, head
	}
	head.Next = nil
	lastNode.Next = head
	return newHead, head, tailNode

}