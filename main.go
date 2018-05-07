package main

import (
	"fmt"

	"github.com/VimleshS/ds_algo/binarysearch"
	"github.com/VimleshS/ds_algo/bubblesort"
	"github.com/VimleshS/ds_algo/fib"
	"github.com/VimleshS/ds_algo/mergesort"
	"github.com/VimleshS/ds_algo/quicksort"
)

func main() {
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
