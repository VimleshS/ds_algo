package binarysearch

type BSort struct{}

func (b BSort) Binarysearch(slice []int, left, right, elem int) bool {
	if left >= right {
		return false
	}

	mid := (left + right) / 2
	if slice[mid] == elem {
		return true
	}
	if slice[mid] > elem {
		return b.Binarysearch(slice, left, mid-1, elem)
	}
	return b.Binarysearch(slice, mid+1, right, elem)
}

//O(logN)
