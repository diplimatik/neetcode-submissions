func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := new(ListNode)
	currHead := ans
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		currHead.Next = &ListNode{Val: sum % 10}
		currHead = currHead.Next
	}
	if carry > 0 {
		currHead.Next = &ListNode{Val: carry}
	}
	return ans.Next
}