package src

import "fmt"

type Heap struct {
	List []int
}

func (h *Heap) Push(v int) {
	h.List = append(h.List, v)

	idx := len(h.List) - 1
	for idx >= 0 {
		pIdx := (idx - 1) / 2
		if pIdx < 0 {
			break //no parent
		}

		//if h.List[idx] > h.List[pIdx] {
		if h.List[idx] < h.List[pIdx] {
			h.List[idx], h.List[pIdx] = h.List[pIdx], h.List[idx]
			idx = pIdx
		} else {
			break
		}
	}
}

func (h *Heap) Pop() int {
	if len(h.List) == 0 {
		return 0
	}

	top := h.List[0]
	last := h.List[len(h.List)-1]
	h.List = h.List[:len(h.List)-1]

	h.List[0] = last
	idx := 0

	for idx < len(h.List) {
		sIdx := -1

		lIdx := idx*2 + 1
		if lIdx >= len(h.List) {
			break // no child
		}
		//if h.List[lIdx] > h.List[idx] {
		if h.List[lIdx] < h.List[idx] {
			sIdx = lIdx
		}

		rIdx := idx*2 + 2
		if rIdx < len(h.List) { // has a child
			//if h.List[rIdx] > h.List[idx] {
			if h.List[rIdx] > h.List[idx] {
				//if sIdx < 0 || h.List[sIdx] < h.List[rIdx] {
				if sIdx < 0 || h.List[sIdx] > h.List[rIdx] {
					sIdx = rIdx
				}
			}
		}

		if sIdx < 0 {
			break
		}

		h.List[idx], h.List[sIdx] = h.List[sIdx], h.List[idx]
		idx = sIdx
	}

	return top
}

func (h *Heap) Print() {
	fmt.Println(h.List)
}

func (h *Heap) Count() int {
	return len(h.List)
}
