package main

import "fmt"

// Node ...
type Node struct {
	Next  *Node
	Value string
}

//LinkedList ...struct so that in case if head changes, all other nodes are informed.
type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Add(node *Node) {
	curNode := l.Head
	if curNode == nil {
		curNode = node
		return
	}
	for curNode.Next != nil {
		curNode = curNode.Next
	}
	curNode.Next = node
}

func (l *LinkedList) Prepend(node *Node) {
	node.Next, l.Head = l.Head, node
}

func (l *LinkedList) PrintNodeValues() {
	curr_node := l.Head
	for curr_node != nil {
		fmt.Println(curr_node.Value)
		curr_node = curr_node.Next
	}
}

func (l *LinkedList) Delete(value string) {
	if l.Head == nil {
		return
	}

	if l.Head.Value == value {
		next := l.Head.Next
		l.Head = next
		return
	}

	curNode := l.Head.Next
	var preNode *Node
	for curNode.Value != value {
		preNode = curNode
		curNode = curNode.Next
	}
	preNode.Next = curNode.Next
}

func testLinkedList() {
	n := Node{Value: "Head"}
	ll := LinkedList{&n}
	ll.Add(&Node{Value: "1"})
	ll.Add(&Node{Value: "11"})
	ll.Add(&Node{Value: "5"})
	ll.Add(&Node{Value: "166"})
	ll.Prepend(&Node{Value: "NewNode"})
	ll.PrintNodeValues()
	ll.Delete("NewNode")
	fmt.Println("--------------------------------------")
	ll.PrintNodeValues()
}
