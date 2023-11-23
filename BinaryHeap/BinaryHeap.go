package binaryheap

type BinaryHeap struct {
	heap []int
}

func NewBinaryHeap() *BinaryHeap {
	return &BinaryHeap{heap: make([]int, 0)}
}

func (h *BinaryHeap) Insert(value int) {
	h.heap = append(h.heap, value)
	h.bubbleUp(len(h.heap) - 1)
}

func (h *BinaryHeap) bubbleUp(index int) {
	parent := (index - 1) / 2
	if parent < 0 {
		return
	}
	if h.heap[parent] > h.heap[index] {
		h.heap[parent], h.heap[index] = h.heap[index], h.heap[parent]
		h.bubbleUp(parent)
	}
}

func (h *BinaryHeap) ExtractMin() int {
	if len(h.heap) == 0 {
		return -1
	}
	min := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.bubbleDown(0)
	return min
}

func (h *BinaryHeap) bubbleDown(index int) {
	left := 2*index + 1
	right := 2*index + 2
	if left >= len(h.heap) {
		return
	}
	if right >= len(h.heap) {
		if h.heap[left] < h.heap[index] {
			h.heap[left], h.heap[index] = h.heap[index], h.heap[left]
		}
		return
	}
	if h.heap[left] < h.heap[right] {
		if h.heap[left] < h.heap[index] {
			h.heap[left], h.heap[index] = h.heap[index], h.heap[left]
			h.bubbleDown(left)
		}
	} else {
		if h.heap[right] < h.heap[index] {
			h.heap[right], h.heap[index] = h.heap[index], h.heap[right]
			h.bubbleDown(right)
		}
	}
}

func (h *BinaryHeap) Peek() int {
	if len(h.heap) == 0 {
		return -1
	}
	return h.heap[0]
}
