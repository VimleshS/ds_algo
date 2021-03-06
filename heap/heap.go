//                (0) 2
//          		/    \
//             (1) 10     65(2)
//                /   \   /  \
//          (3)	11    14  88
//                    (4) (5)
//
// Array       	2, 10, 65, 11, 14, 88
// Index       	0  1   2   3   4   5
//
//  Calc       	Lchild =  2i + 1
//          	RChild =  2i + 2
//          	Parent =  (i-1) / 2

package heap

type heap struct {
	store []int
}

func NewHeap() *heap {
	return &heap{store: []int{}}
}

func (h *heap) Add(v int) {
	h.store = append(h.store, v)
	h.heapifyUp()
}

func (h *heap) Peek() int {
	retVal := h.store[0]
	h.store[0] = h.store[len(h.store)-1]
	h.store = h.store[0 : len(h.store)-1]
	h.heapifyDown()
	return retVal
}

func (h *heap) heapifyDown() {
	r := 0
	lchild := getLeftChild(r)
	for h.hasNode(lchild) {
		smallestIdx := lchild

		if rchild := getRigthChild(r); h.hasNode(rchild) && h.store[rchild] < h.store[lchild] {
			smallestIdx = rchild
		}

		if h.store[r] < h.store[smallestIdx] {
			break
		} else {
			h.store[r], h.store[smallestIdx] = h.store[smallestIdx], h.store[r]
			r = smallestIdx
			lchild = getLeftChild(r)
		}
	}
}

func (h *heap) heapifyUp() {
	i := len(h.store) - 1
	parentI := getParentIdx(i)
	for h.hasNode(parentI) && (h.store[parentI] > h.store[i]) {
		//swap
		h.store[parentI], h.store[i] = h.store[i], h.store[parentI]
		i = parentI
		parentI = getParentIdx(i)
	}
}

func (h *heap) hasNode(idx int) bool {
	if (idx > len(h.store)-1) || (idx < 0) {
		return false
	}
	return true
}

func getParentIdx(idx int) int {
	return (idx - 1) / 2
}
func getLeftChild(idx int) int {
	return 2*idx + 1
}

func getRigthChild(idx int) int {
	return 2*idx + 2
}
