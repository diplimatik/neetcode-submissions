/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	visitedNodes := make(map[*ListNode]int)
	for head != nil {
		if visitedNodes[head] != 0 {
			return true
		}
		visitedNodes[head] = 1
		head = head.Next
	}
	return false
}

