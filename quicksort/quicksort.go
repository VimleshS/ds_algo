package quicksort

type QC struct{}

func (qc *QC) Quicksort(slice []int, left, right int) {
	if left >= right {
		return
	}
	pivot := slice[(left+right)/2]
	index := qc.partition(slice, left, right, pivot)
	qc.Quicksort(slice, left, index-1)
	qc.Quicksort(slice, index, right)
}

func (qc *QC) partition(slice []int, left, right, pivot int) int {
	for left <= right {
		for slice[left] < pivot {
			left++
		}

		for slice[right] > pivot {
			right--
		}

		if left <= right {
			slice[left], slice[right] = slice[right], slice[left]
			left++
			right--
		}

	}
	return left
}
