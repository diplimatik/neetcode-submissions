type ListNode struct {
	Key  int
	Val  int
	Next *ListNode
}
type LRUCache struct {
	freeSpace  int
	mapStorage map[int]*ListNode
	head       *ListNode
	tail       *ListNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		freeSpace:  capacity,
		mapStorage: make(map[int]*ListNode, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if prevNode, ok := this.mapStorage[key]; !ok {
		return -1
	} else {
		if prevNode == nil {
			return this.head.Val
		}
		currNode := prevNode.Next
		prevNode.Next = currNode.Next
		this.mapStorage[this.head.Key] = currNode
		this.mapStorage[currNode.Key] = nil
		if currNode.Next != nil {
			this.mapStorage[currNode.Next.Key] = prevNode
		}
		currNode.Next = this.head
		this.head = currNode
		if currNode == this.tail {
			this.tail = prevNode
		}
		return currNode.Val
	}
}

func (this *LRUCache) Put(key int, value int) {
	newNode := &ListNode{
		key,
		value,
		this.head,
	}
	
	if _, ok := this.mapStorage[key]; ok {
		this.Get(key)
		this.head.Val = value
		return
	}
	
	if this.freeSpace > 0 {
		if this.tail == nil {
			this.tail = newNode
		}
		this.freeSpace--
	} else {
		newTail := this.mapStorage[this.tail.Key]
		if newTail != nil {
			if newTail.Key != newNode.Key {
				delete(this.mapStorage, this.tail.Key)
				newTail.Next = nil
			}
			this.tail = newTail
		} else if this.tail == this.head {
			delete(this.mapStorage, this.tail.Key)
			this.tail = newNode
			this.head = nil
		}
	}

	if this.head != nil {
		this.mapStorage[this.head.Key] = newNode
	}
	this.head = newNode
	this.mapStorage[key] = nil
}