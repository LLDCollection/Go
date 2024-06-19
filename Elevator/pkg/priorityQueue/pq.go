package priorityqueue

type MinIntHeap []int

func (h MinIntHeap) Len() int {
	return len(h)
}

func (h MinIntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinIntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinIntHeap) Pop() interface{} {
	old := *h
	item := old[len(old) - 1]
	*h = old[0: len(old) - 1]
	return item
}
