package p2558

import "math"

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

func (h *Heap[T]) UpdateTop(fn func(v T) T) T {
	v := fn(h.vals[0])
	h.vals[0] = v
	if !h.down(0, len(h.vals)) {
		h.up(0)
	}
	return v
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

func pickGifts(gifts []int, k int) int64 {
	h := NewHeap(gifts, func(a, b int) bool {
		return a > b
	})
	for i := 0; i < k; i++ {
		h.UpdateTop(func(v int) int {
			return int(math.Sqrt(float64(v)))
		})
	}
	var sum int64
	for _, val := range gifts {
		sum += int64(val)
	}
	return sum
}
