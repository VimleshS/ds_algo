package bubblesort

import "fmt"

func BSort(slice []int) bool {
	//Base Case
	sorted := true

	//This will progressively pop element up the array
	for index := 0; index < len(slice)-1; index++ {
		if slice[index] > slice[index+1] {
			sorted = false
			slice[index], slice[index+1] = slice[index+1], slice[index]
			//break // this makes loop inefficent
		}
	}

	//Recursive loop, which pushes to call stack
	for sorted == false {
		sorted = BSort(slice)
	}

	//Returns from a call stack
	return sorted
}

func BSort_2(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if slice[i] < slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	fmt.Println(slice)
}
