package main

import "fmt"

func testTrees() {
	t := Tree{}
	t.add(2)
	t.add(100)
	t.add(23)
	t.add(6)
	t.add(50)
	fmt.Println("\n------In---ASC---")
	t.traverse()
	fmt.Println("\n------Pr--------")
	t.traversePre()
	fmt.Println("\n------Po--------")
	t.traversePost()
	fmt.Println("\n------In---DESC---")
	t.traverseDesc()
	fmt.Println()
}

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func (t *Tree) add(value int) {
	if t.Value > value {
		if t.Left == nil {
			t.Left = &Tree{Value: value}
		} else {
			t.Left.add(value)
		}
	} else {
		if t.Right == nil {
			t.Right = &Tree{Value: value}
		} else {
			t.Right.add(value)
		}
	}
}

func (t *Tree) traverse() {
	//InOrder
	if t.Left != nil {
		t.Left.traverse()
	}
	fmt.Printf("%v \t ", t.Value)
	if t.Right != nil {
		t.Right.traverse()
	}
}

func (t *Tree) traversePre() {
	fmt.Printf("%v \t ", t.Value)
	if t.Left != nil {
		t.Left.traversePre()
	}

	if t.Right != nil {
		t.Right.traversePre()
	}
}

func (t *Tree) traversePost() {
	if t.Left != nil {
		t.Left.traversePost()
	}
	if t.Right != nil {
		t.Right.traversePost()
	}
	fmt.Printf("%v \t ", t.Value)
}

func (t *Tree) traverseDesc() {
	if t.Right != nil {
		t.Right.traverseDesc()
	}
	fmt.Printf("%v \t ", t.Value)
	if t.Left != nil {
		t.Left.traverseDesc()
	}
}
