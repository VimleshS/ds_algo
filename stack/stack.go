package stack

import "fmt"

type Node struct {
	data int
	prev *Node
}
type Stack struct {
	top *Node
}

func (s *Stack) Push(data int) {
	d := &Node{data: data}
	d.prev = s.top
	s.top = d
}

//instead of err return -1 for not found
func (s *Stack) Pop() int {
	if s.top == nil {
		return -1
	}

	retVal := s.top.data
	s.top = s.top.prev
	return retVal
}

func (s *Stack) Print() {
	t := s.top
	for t != nil {
		fmt.Println(t.data)
		t = t.prev
	}
}
