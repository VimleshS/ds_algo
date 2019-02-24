package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) Enqueue(data int) {
	n := &Node{data: data}
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
}
func (l *LinkedList) Dequeue() (int, error) {
	if l.head == nil {
		return -1, fmt.Errorf("Queue Empty")
	}
	retVal := l.head.data
	l.head = l.head.next
	return retVal, nil
}

func (l *LinkedList) Print() {
	if l.head == nil {
		return
	}
	t := l.head
	for t != nil {
		fmt.Println(t.data)
		t = t.next
	}
}
