func copyRandomList(head *Node) *Node {
	dummy := new(Node)
	nodes := make(map[*Node]*Node)
	for currHead, currDummy := head, dummy; currHead != nil; currHead = currHead.Next {
		var currRandCpy *Node
		var ok bool
		if currRandCpy, ok = nodes[currHead.Random]; !ok && currHead.Random != nil {
			currRandCpy = &Node{Val: currHead.Random.Val}
			nodes[currHead.Random] = currRandCpy
		}
		if currHeadCopy, ok := nodes[currHead]; ok {
			currHeadCopy.Random = currRandCpy
			currDummy.Next = currHeadCopy
		} else {
			currHeadCopy = &Node{Val: currHead.Val, Random: currRandCpy}
			currDummy.Next = currHeadCopy
			nodes[currHead] = currHeadCopy
		}
		currDummy = currDummy.Next
	}
	return dummy.Next
}