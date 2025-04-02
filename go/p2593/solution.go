package p2593

// container/heap is a heap of garbage amirite!?
type Heap[T any] struct {
	vals []T
	less func(a, b T) bool
}

func (h *Heap[T]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(h.vals[j], h.vals[i]) {
			break
		}
		h.vals[i], h.vals[j] = h.vals[j], h.vals[i]
		j = i
	}
}

func (h *Heap[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(h.vals[j2], h.vals[j1]) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.less(h.vals[j], h.vals[i]) {
			break
		}
		h.vals[i], h.vals[j] = h.vals[j], h.vals[i]
		i = j
	}
	return i > i0
}

func NewHeap[T any](vals []T, less func(a, b T) bool) *Heap[T] {
	h := &Heap[T]{
		vals: vals,
		less: less,
	}

	n := len(vals)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}

	return h
}

func (h *Heap[T]) Pop() T {
	n := len(h.vals) - 1
	h.vals[0], h.vals[n] = h.vals[n], h.vals[0]
	h.down(0, n)

	v := h.vals[n]
	h.vals = h.vals[:n]
	return v
}

func (h *Heap[T]) Len() int {
	return len(h.vals)
}

func findScore(nums []int) int64 {
	n := len(nums)
	idxs := make([]int, n)
	for i := range nums {
		idxs[i] = i
	}
	marked := make([]bool, n)
	h := NewHeap(idxs, func(i, j int) bool {
		iv := nums[i]
		jv := nums[j]
		if iv == jv {
			return i < j
		}
		return iv < jv
	})

	var score int64
	for h.Len() > 0 {
		idx := h.Pop()
		if marked[idx] {
			continue
		}
		score += int64(nums[idx])
		marked[idx] = true
		if pi := idx - 1; pi >= 0 {
			marked[pi] = true
		}
		if ni := idx + 1; ni < n {
			marked[ni] = true
		}
	}
	return score
}
