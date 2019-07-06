package src

import "fmt"

type Heap struct {
	list []int
}

// 맨 뒤에 붙인휘, 부모 노드와 비교 하며 올라간다.
func (h *Heap) Push(v int) {
	h.list = append(h.list, v)

	idx := len(h.list) - 1
	for idx >= 0 {
		parentIdx := (idx - 1) / 2
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

// 맨 뒤에 값을 일단 root로 보내고,
// 자식으로 내려가면서 비교
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
		lidx := idx*2 + 1
		// no child
		if lidx >= len(h.list) {
			break
		}

		if h.list[lidx] > h.list[idx] {
			swapIdx = lidx
		}

		ridx := idx*2 + 2
		// has child
		if ridx < len(h.list) {
			if h.list[ridx] > h.list[idx] {
				if swapIdx < 0 || h.list[swapIdx] < h.list[ridx] {
					swapIdx = ridx
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
