/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	dummyHead := dummy

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			dummyHead.Next = list2
			list2 = list2.Next
		} else {
			dummyHead.Next = list1
			list1 = list1.Next
		}
		dummyHead = dummyHead.Next
	}
	if list1 != nil {
		dummyHead.Next = list1
	} else {
		dummyHead.Next = list2
	}
	return dummy.Next
}