package core

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	a := []int{2, 3, 4, 1, 5}

	heap := NewNodeHeap()

	for i := range a {
		heap.Push(&Node{
			Cost: a[i],
		})
	}

	for heap.Len() > 0 {
		fmt.Printf("%v\n", heap.Pop())
	}
}
