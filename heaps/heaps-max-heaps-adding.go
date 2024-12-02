package heaps

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(value int) {
	h.array = append(h.array, value)
	h.maxHeapifyUp(len(h.array) - 1)
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return (2*index + 1)
}

func right(index int) int {
	return (2*index + 2)
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
func (h *MaxHeap) maxHeapifyUp(index int) {

	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)

		index = parent(index)
	}
}

func main() {
	h := MaxHeap{}

	h.Insert(10)
	h.Insert(15)
	h.Insert(20)
	h.Insert(5)
	h.Insert(30)
	h.Insert(25)

	fmt.Println("Heap após inserções:", h.array)
}
