package heap

import (
	"errors"
	"fmt"
	"strings"
)

// Heap is a priority queue data structure.
type Heap struct {
	Container []any
	Comp      func(child any, parent any) bool
}

// New return a pointer to an empty heap
func New() *Heap {
	return &Heap{}
}

// AddComparator sets the comparator function used for comparision during heapify operation
func (h *Heap) AddComparator(Comp func(child any, parent any) bool) {
	h.Comp = Comp
}

// Len returns length of heap
func (h *Heap) Len() int {
	return len(h.Container)
}

// Cap return capacity of Heap
func (h *Heap) Cap() int {
	return cap(h.Container)
}

// Push inserts the given item in heap
func (h *Heap) Push(item any) error {
	if h.Empty() == false && strings.EqualFold(fmt.Sprintf("%T", item), fmt.Sprintf("%T", h.Container[0])) == false {
		return errors.New(fmt.Sprintf("expected type %s, got %s\n", fmt.Sprintf("%T", item), fmt.Sprintf("%T", h.Container[0])))
	}
	h.Container = append(h.Container, item)
	child := h.Len() - 1
	for child > 0 {
		parent := (child - 1) / 2
		if h.Comp(h.Container[child], h.Container[parent]) {
			break
		}
		h.Container[parent], h.Container[child] = h.Container[child], h.Container[parent]
		child = parent
	}
	fmt.Println(h.Container)
	return nil
}

// Empty checks whether the queue is empty
func (h *Heap) Empty() bool {
	return h.Len() == 0
}

// Return the item with highest priority
func (h *Heap) Seek() (any, error) {
	if h.Len() == 0 {
		return 0, errors.New("empty container")
	}
	return h.Container[0], nil
}

// Pop removes the item with highest priority
func (h *Heap) Pop() {
	parent := 0
	h.Container[parent] = h.Container[h.Len()-1]
	for parent < h.Len()-1 {
		child1 := 2*parent + 1
		child2 := 2*parent + 2
		if child1 >= h.Len()-1 {
			break
		}
		if child2 >= h.Len()-1 {
			if h.Comp(h.Container[child1], h.Container[parent]) == false {
				h.Container[parent], h.Container[child1] = h.Container[child1], h.Container[parent]
				parent = child1
			} else {
				break
			}
		} else if h.Comp(h.Container[child1], h.Container[parent]) == false && h.Comp(h.Container[child1], h.Container[child2]) == false {
			h.Container[parent], h.Container[child1] = h.Container[child1], h.Container[parent]
			parent = child1
		} else if h.Comp(h.Container[child2], h.Container[parent]) == false && h.Comp(h.Container[child2], h.Container[child1]) == false {
			h.Container[parent], h.Container[child2] = h.Container[child2], h.Container[parent]
			parent = child2
		} else {
			break
		}

	}
	h.Container = h.Container[0 : h.Len()-1]
}
