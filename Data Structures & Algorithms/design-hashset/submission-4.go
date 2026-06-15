type LinkedList struct {
	next *LinkedList
	item int
}

type MyHashSet struct {
	buckets []*LinkedList
}

func Constructor() MyHashSet {
	return MyHashSet{
		buckets: make([]*LinkedList, 32),
	}
}

func (this *MyHashSet) Add(key int) {
	if this.buckets[key%len(this.buckets)] == nil {
		this.buckets[key%len(this.buckets)] = &LinkedList{item: key}
		return
	}

	currItem := this.buckets[key%len(this.buckets)]
	for {
		if currItem.item == key {
			return
		}
		if currItem.next == nil {
			break
		}
		currItem = currItem.next
	}
	currItem.next = &LinkedList{item: key}
}

func (this *MyHashSet) Remove(key int) {
	currItem := this.buckets[key%len(this.buckets)]
	if currItem == nil {
		return
	}
	if currItem.item == key {
		this.buckets[key%len(this.buckets)] = this.buckets[key%len(this.buckets)].next
		return
	}
	for currItem.next != nil {
		if currItem.next.item == key {
			currItem.next = currItem.next.next
			return
		}
		currItem = currItem.next
	}
}

func (this *MyHashSet) Contains(key int) bool {
	if this.buckets[key%len(this.buckets)] == nil {
		return false
	}
	currItem := this.buckets[key%len(this.buckets)]
	if currItem.item == key {
		return true
	}
	for currItem != nil {
		if currItem.item == key {
			return true
		}
		currItem = currItem.next
	}

	return false
}
