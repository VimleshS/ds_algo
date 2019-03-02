package circularlinkedlist

//https://www.pythoncentral.io/circular-queue/

import "fmt"

type circularLinkedlist struct {
	store   []int
	front   int
	tail    int
	maxsize int
}

//No of element == Maxsize -1
func NewCircularList(maxsise int) *circularLinkedlist {
	return &circularLinkedlist{maxsize: maxsise}
}

func (c *circularLinkedlist) Enqueue(v int) {
	if c.Size() == c.maxsize-1 {
		fmt.Printf("cannot add %v size is full \n", v)
		return
	}

	c.store = append(c.store, v)
	c.tail = (c.tail + 1) % c.maxsize
}

func (c *circularLinkedlist) Dequeue() int {
	if c.Size() <= 0 {
		fmt.Println("No more element to pop")
		return -1
	}

	retVal := c.store[c.front]
	c.front = (c.front + 1) % c.maxsize
	return retVal
}

//Actual filled size
func (c *circularLinkedlist) Size() int {
	if c.tail >= c.front {
		return c.tail - c.front
	}
	return c.maxsize - (c.front - c.tail)
}
