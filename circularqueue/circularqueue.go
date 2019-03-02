package circularqueue

//https://www.pythoncentral.io/circular-queue/

import "fmt"

type circularqueue struct {
	store   []int
	front   int
	tail    int
	maxsize int
}

//No of element == Maxsize -1
func NewCircularQueue(maxsize int) *circularqueue {
	return &circularqueue{
		store:   make([]int, maxsize),
		maxsize: maxsize,
	}
}

func (c *circularqueue) Enqueue(v int) {
	if c.Size() == c.maxsize-1 {
		fmt.Printf("cannot add %v size is full \n", v)
		return
	}

	//c.store = append(c.store, v)
	c.store[c.tail] = v
	c.tail = (c.tail + 1) % c.maxsize
}

func (c *circularqueue) Dequeue() int {
	if c.Size() <= 0 {
		fmt.Println("No more element to pop")
		return -1
	}

	retVal := c.store[c.front]
	c.front = (c.front + 1) % c.maxsize
	return retVal
}

//Actual filled size
func (c *circularqueue) Size() int {
	if c.tail >= c.front {
		return c.tail - c.front
	}
	return c.maxsize - (c.front - c.tail)
}
