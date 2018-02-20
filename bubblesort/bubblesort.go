package bubblesort

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
