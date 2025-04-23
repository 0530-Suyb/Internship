package main

type LinkedNode struct {
	val  int
	next *LinkedNode
}

type MyLinkedList struct {
	dummyHead *LinkedNode
	size      int
}

func Constructor() MyLinkedList {
	return MyLinkedList{dummyHead: &LinkedNode{}}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index < 0 {
		index = 0
	}
	this.size++
	cur := this.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	newNode := &LinkedNode{val: val, next: cur.next}
	cur.next = newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	this.size--
	cur := this.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.next = cur.next.next
}
