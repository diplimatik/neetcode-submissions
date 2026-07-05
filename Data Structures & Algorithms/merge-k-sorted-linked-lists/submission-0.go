/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		l1 := lists[0]
		l2 := lists[1]

		if l1 == nil && l2 == nil {
			lists = lists[1:]
			continue
		} else if l1 == nil {
			lists = lists[1:]
			continue
		} else if l2 == nil {
			lists = append(lists, l1)
			lists = lists[2:]
			continue
		}

		lists = lists[2:]
		var minVal int
		if l1.Val > l2.Val {
			minVal = l2.Val
			l2 = l2.Next
		} else {
			minVal = l1.Val
			l1 = l1.Next
		}

		sortedList := &ListNode{Val: minVal}
		currNode := sortedList

		for l1 != nil || l2 != nil {
			if l1 != nil && l2 != nil {
				if l1.Val < l2.Val {
					currNode.Next = &ListNode{Val: l1.Val}
					currNode = currNode.Next
					l1 = l1.Next
				} else {
					currNode.Next = &ListNode{Val: l2.Val}
					currNode = currNode.Next
					l2 = l2.Next
				}
			} else if l1 != nil {
				for l1 != nil {
					currNode.Next = &ListNode{Val: l1.Val}
					currNode = currNode.Next
					l1 = l1.Next
				}
			} else {
				for l2 != nil {
					currNode.Next = &ListNode{Val: l2.Val}
					currNode = currNode.Next
					l2 = l2.Next
				}
			}
		}

		lists = append(lists, sortedList)
	}

	return lists[0]
}
