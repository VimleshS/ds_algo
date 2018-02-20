package mergesort

func MergeSort(slice []int) []int {
	if len(slice) == 1 {
		return slice
	}
	mid := len(slice) / 2
	m1 := MergeSort(slice[:mid])
	m2 := MergeSort(slice[mid:])
	return merge(m1, m2)
}

func merge(m1, m2 []int) []int {
	_t := []int{}

	ri, li := 0, 0
	for len(m1) > li && len(m2) > ri {
		if m1[li] >= m2[ri] {
			_t = append(_t, m2[ri])
			ri++
		} else {
			_t = append(_t, m1[li])
			li++
		}
	}

	if len(m1) > li {
		_t = append(_t, m1[li:]...)
	}
	if len(m2) > ri {
		_t = append(_t, m2[ri:]...)
	}
	return _t
}
