package main

import (
	"fmt"

	"github.com/VimleshS/ds_algo/graph"

	"github.com/VimleshS/ds_algo/binarysearch"
	"github.com/VimleshS/ds_algo/bubblesort"
	"github.com/VimleshS/ds_algo/circularqueue"
	"github.com/VimleshS/ds_algo/fib"
	Vheap "github.com/VimleshS/ds_algo/heap"
	Iheap "github.com/VimleshS/ds_algo/intheap"
	"github.com/VimleshS/ds_algo/linkedlist"
	"github.com/VimleshS/ds_algo/mergesort"
	"github.com/VimleshS/ds_algo/quicksort"
	"github.com/VimleshS/ds_algo/stack"
)

func main() {
	clist := circularqueue.NewCircularQueue(5)
	clist.Enqueue(1)
	clist.Enqueue(12)
	clist.Enqueue(5)
	clist.Enqueue(122)
	clist.Enqueue(232)
	fmt.Println(clist.Dequeue())
	// clist.Enqueue(232)
	// clist.Enqueue(89)
	// fmt.Println(clist.Dequeue())
	// clist.Enqueue(78)
	// clist.Enqueue(189)
	// clist.Enqueue(178)
	fmt.Println(clist.Dequeue())
	fmt.Println(clist.Dequeue())
	fmt.Println(clist.Dequeue())
	fmt.Println(clist.Dequeue())
	fmt.Println(clist.Dequeue())
	fmt.Println(clist.Dequeue())

	return

	sta := stack.Stack{}
	sta.Push(3)
	sta.Push(2)
	sta.Push(7)
	sta.Push(9)
	sta.Print()
	fmt.Println("--->", sta.Pop())
	sta.Push(9993)
	sta.Print()

	return

	l := &linkedlist.LinkedList{}
	l.Enqueue(3)
	l.Enqueue(13)
	l.Enqueue(31)
	l.Enqueue(113)
	l.Enqueue(30)
	l.Print()
	p, _ := l.Dequeue()
	fmt.Println("de---", p)
	l.Enqueue(111130)
	l.Print()
	return
	h := Vheap.NewHeap()
	h.Add(34)
	h.Add(1)
	fmt.Println(h.Peek())
	fmt.Println(h)

	minHeap := Iheap.NewHeap(Iheap.MIN)
	minHeap.Add(34)
	minHeap.Add(3)
	minHeap.Add(31)
	minHeap.Add(30)
	minHeap.Add(20)
	minHeap.Add(1)
	fmt.Println(h)
	fmt.Println(minHeap.Peek())
	fmt.Println(minHeap.Peek())
	fmt.Println(minHeap.Peek())
	fmt.Println(minHeap.Peek())
	fmt.Println(minHeap.Peek())
	fmt.Println(minHeap.Peek())
	fmt.Println(minHeap)
	maxHeap := Iheap.NewHeap(Iheap.MAX)
	maxHeap.Add(34)
	maxHeap.Add(3)
	maxHeap.Add(31)
	maxHeap.Add(30)
	maxHeap.Add(20)
	maxHeap.Add(1)
	fmt.Println(h)
	fmt.Println(maxHeap.Peek())
	fmt.Println(maxHeap.Peek())
	fmt.Println(maxHeap.Peek())
	fmt.Println(maxHeap.Peek())
	fmt.Println(maxHeap.Peek())
	fmt.Println(maxHeap.Peek())
	fmt.Println(maxHeap)
	return

	graph.TopologicalSort()

	return

	testTrees()
	return
	q := Queue{}
	q.Add("Hello")
	q.Add("Vimlesh")
	q.Add("Sharma")
	q.PrintQueue()
	return

	testLinkedList()
	return
	bs := []int{150, 40, 6, 2, 7, 10, 100, 2}
	fmt.Println(mergesort.MergeSort(bs))
	return

	bubblesort.BSort(bs)
	fmt.Println(bs)
	return

	// fmt.Println(fib.Fib(5))
	fmt.Println(fib.FibEfficient(50))
	return

	s := []int{1, 4, 6, 2, 5, 10, 3}
	fmt.Println(binarysearch.BSort{}.Binarysearch(s, 0, len(s)-1, -1))
	return

	qc := quicksort.QC{}
	qc.Quicksort(s, 0, len(s)-1)
	fmt.Println(s)
}
