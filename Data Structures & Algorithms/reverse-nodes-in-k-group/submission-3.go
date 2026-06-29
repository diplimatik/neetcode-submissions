func reverseKGroup(head *ListNode, k int) *ListNode {
	prev := head
	tail := head
	ans := head
	ansFlag := false
	for {
		if kCheck(prev, k) {
			curr := prev.Next
			for c := k; c > 1; c-- {
				tmp := curr.Next
				curr.Next = prev
				prev = curr
				curr = tmp
			}
			if !ansFlag {
				ans = prev
				ansFlag = true
			} else {
				tail.Next = prev
				tail = head
			}

			head.Next = curr
			head = curr
			prev = curr
		} else {
			return ans
		}
	}
}

func kCheck(head *ListNode, k int) bool {
	curr := head
	for curr != nil && k > 0 {
		curr = curr.Next
		k--
	}
	if k > 0 {
		return false
	}
	return true
}