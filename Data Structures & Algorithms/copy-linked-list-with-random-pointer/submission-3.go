/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	dummy := new(Node)
	nodes := make(map[*Node]*Node)
	for currHead, currDummy := head, dummy; currHead != nil; currHead = currHead.Next {
		currDummy.Next = &Node{Val: currHead.Val}
		nodes[currHead] = currDummy.Next
		currDummy = currDummy.Next
	}
	for currHead, currDummy := head, dummy.Next; currDummy != nil; currDummy = currDummy.Next {
		currDummy.Random = nodes[currHead.Random]
		currHead = currHead.Next
	}
	return dummy.Next
}