func reorderList(head *ListNode) {
	nodeCount := 0
	for countHead := head; countHead != nil; countHead = countHead.Next {
		nodeCount++
	}
	if nodeCount < 3 {
		return
	}
	prev := head
	nodeCount /= 2
	for nodeCount > 1 {
		prev = prev.Next
		nodeCount--
	}
	tail := prev
	prev = prev.Next
	tail.Next = nil
	curr := prev.Next
	prev.Next = nil
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	count := 0
	for prev != nil && head != nil {
		if count%2 == 0 {
			tmp := head.Next
			head.Next = prev
			head = tmp
		} else {
			tmp := prev.Next
			prev.Next = head
			prev = tmp
		}
		count++
	}
}