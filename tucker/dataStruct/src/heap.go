package src

import "fmt"

type Heap struct {
	list []int
}

func (h *Heap) Push(v int) {
	h.list = append(h.list, v)

	idx := len(h.list) - 1
	for idx >= 0 {
		parentIdx := (idx - 1) / 2 // 부모노드
		if parentIdx < 0 {
			break
		}

		if h.list[idx] > h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *Heap) Print() {
	fmt.Println(h.list)
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
		si := -1
		li := idx*2 + 1        // left index
		if li >= len(h.list) { // no child node
			break
		}
		if h.list[li] > h.list[idx] {
			si = li
		}

		ri := idx*2 + 2 // right index
		if ri < len(h.list) {
			if h.list[ri] > h.list[idx] {
				if si < 0 || h.list[si] < h.list[ri] {
					si = ri
				}
			}
		}

		if si < 0 {
			break
		}

		h.list[idx], h.list[si] = h.list[si], h.list[idx]
		idx = si
	}

	return top
}
