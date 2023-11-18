package circularbuffer

import "fmt"

type CircularBuffer struct {
	data    []int
	currIdx int
}

func NewCircularBuffer() *CircularBuffer {
	cb := &CircularBuffer{data: make([]int, 10), currIdx: 0}
	return cb
}

// Insert inserts a value into the circular buffer and if buffer is full,
// the first value is removed to make room for the new value.
func (cb *CircularBuffer) Insert(value int) *CircularBuffer {
	if (cb.currIdx + 1) > len(cb.data) {
		cb.currIdx = 0
	}
	cb.data[cb.currIdx] = value
	cb.currIdx++
	return cb
}

func (cb *CircularBuffer) Remove(value int) *CircularBuffer {
	if cb.currIdx == 0 {
		return cb
	}

	foundIdx := -1
	for i := 0; i < cb.currIdx; i++ {
		if cb.data[i] == value {
			foundIdx = i
			break
		}
	}

	if foundIdx == -1 {
		return cb
	}

	for i := foundIdx; i < cb.currIdx-1; i++ {
		cb.data[i] = cb.data[i+1]
	}
	cb.currIdx--
	cb.data[cb.currIdx] = 0

	return cb
}

func (cb *CircularBuffer) Print() string {
	var result string
	result += "["
	for i := 0; i < cb.currIdx; i++ {
		if i == cb.currIdx-1 {
			result += fmt.Sprintf("%d", cb.data[i])
		} else {
			result += fmt.Sprintf("%d, ", cb.data[i])
		}
	}
	result += "]"
	return result
}
