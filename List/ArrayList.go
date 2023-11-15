package list

import "fmt"

type ArrayList struct {
	data    []int
	currIdx int
}

func NewArrayList() *ArrayList {
	return &ArrayList{}
}

func (l *ArrayList) Insert(value int) *ArrayList {
	if (l.currIdx + 1) > len(l.data) {
		newData := make([]int, len(l.data)*2)
		copy(newData, l.data)
		l.data = newData
	}
	l.data[l.currIdx] = value
	l.currIdx++
	return l
}

func (l *ArrayList) Remove(value int) *ArrayList {
	if l.currIdx == 0 {
		return l
	}

	foundIdx := -1
	for i := 0; i < l.currIdx; i++ {
		if l.data[i] == value {
			foundIdx = i
			break
		}
	}

	if foundIdx == -1 {
		return l
	}

	for i := foundIdx; i < l.currIdx-1; i++ {
		l.data[i] = l.data[i+1]
	}
	l.currIdx--
	l.data[l.currIdx] = 0

	return l
}

func (l *ArrayList) Print() string {
	var result string
	for i := 0; i < l.currIdx; i++ {
		result += fmt.Sprintf("%d\n", l.data[i])
	}
	return result
}
