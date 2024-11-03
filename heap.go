package main

type Heap struct {
	data []int
	less func(int, int) bool
}

func New(compareFn func(int, int) bool) *Heap {
	return &Heap{
		data: make([]int, 1),
		less: compareFn,
	}
}

func (h *Heap) Push(value int) {
	h.data = append(h.data, value)
	idx := len(h.data) - 1
	parentIdx := idx / 2
	if parentIdx == 0 {
		return
	}
	if h.less(h.data[idx], h.data[parentIdx]) {
		swap(h.data, idx, parentIdx)
	}
}

func (h *Heap) Pop() int {
	val := h.data[1]
	h.data[1] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.rebalance(1)
	return val
}

func (h *Heap) rebalance(idx int) {
	if idx >= len(h.data) || left(idx) >= len(h.data) {
		return
	}
	this := h.data[idx]
	leftVal := h.data[left(idx)]

	if right(idx) >= len(h.data) {
		if h.less(leftVal, this) {
			swap(h.data, idx, left(idx))
		}
		return
	}

	rightVal := h.data[right(idx)]
	if h.less(leftVal, this) || h.less(rightVal, this) {
		if h.less(leftVal, rightVal) {
			swap(h.data, idx, left(idx))
			h.rebalance(left(idx))
		} else {
			swap(h.data, idx, right(idx))
			h.rebalance(right(idx))
		}
	}
}

func swap(data []int, i int, j int) {
	tmp := data[i]
	data[i] = data[j]
	data[j] = tmp
}

func left(idx int) int {
	return idx * 2
}

func right(idx int) int {
	return idx*2 + 1
}
