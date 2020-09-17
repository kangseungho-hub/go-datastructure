package datastructure

import "fmt"

type Heap struct {
	list []int
}

func (h *Heap) Push(data int) {
	h.list = append(h.list, data)

	idx := len(h.list) - 1

	for idx > 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}

		if h.list[parentIdx] < h.list[idx] {
			h.list[parentIdx], h.list[idx] = h.list[idx], h.list[parentIdx]
			idx = parentIdx
		} else {
			break
		}
	}

}

func (h *Heap) Pop() int {
	if len(h.list) == 0 {
		fmt.Println("empty!")
		return 0
	}

	top := h.list[0]

	tail := h.list[len(h.list)-1]
	h.list[0] = tail
	h.list = h.list[:len(h.list)-1]

	idx := 0

	for idx < len(h.list) {

		swapIdx := 0
		leftIdx := idx*2 + 1
		rightIdx := idx*2 + 2

		if leftIdx < len(h.list) {
			if h.list[leftIdx] > h.list[idx] {

				swapIdx = leftIdx
			}

		}

		if rightIdx < len(h.list) {
			if h.list[rightIdx] > h.list[idx] {
				if h.list[swapIdx] < h.list[rightIdx] {
					swapIdx = rightIdx
				}

			}

		}

		if swapIdx == 0 {
			break
		}
		h.list[swapIdx], h.list[idx] = h.list[idx], h.list[swapIdx]
		idx = swapIdx
	}

	return top
}

func (h *Heap) PrintHeap() {
	for i := 0; i < len(h.list); i++ {
		fmt.Printf("%d -> ", h.list[i])
	}
	fmt.Println("")
}
