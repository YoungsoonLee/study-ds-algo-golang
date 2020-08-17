package datastruct

import "fmt"

// 자식 :
// 	left : 2*idx + 1
//	right : 2*idx + 2
// 부모 :
//		(N-1) / 2
//
type Heap struct {
	list []int
}

func (h *Heap) Push(v int) {
	h.list = append(h.list, v)

	idx := len(h.list) - 1

	for idx > 0 {
		pIdx := (idx - 1) / 2

		if pIdx < 0 {
			break
		}
		if h.list[idx] > h.list[pIdx] {
			h.list[idx], h.list[pIdx] = h.list[pIdx], h.list[idx]
			idx = pIdx
		} else {
			break
		}
	}

}

func (h *Heap) Pop() int {
	if len(h.list) == 0 {
		return 0
	}

	top := h.list[0]
	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	h.list[0] = last
	idx := 0
	for idx < len(h.list) {
		swapIdx := -1
		lIdx := (2 * idx) + 1
		if lIdx >= len(h.list) {
			// no child
			break
		}
		if h.list[lIdx] > h.list[idx] {
			swapIdx = lIdx
		}

		rIdx := (2 * idx) + 2
		if rIdx < len(h.list) {
			if h.list[rIdx] > h.list[idx] {
				if swapIdx < 0 || h.list[swapIdx] < h.list[rIdx] {
					swapIdx = rIdx
				}
			}
		}

		if swapIdx < 0 {
			break
		}

		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		idx = swapIdx
	}

	return top

}

func (h *Heap) Print() {
	fmt.Println(h.list)
}
