package p2558

import (
	"container/heap"
	"math"
)

type queue []int

func (q queue) Len() int {
	return len(q)
}

func (q queue) Less(i, j int) bool {
	return q[i] > q[j]
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *queue) Push(x any) {
	*q = append(*q, x.(int))
}

func (q *queue) Pop() any {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

func pickGifts(gifts []int, k int) int64 {
	q := queue(gifts)
	heap.Init(&q)
	for i := 0; i < k; i++ {
		v := heap.Pop(&q).(int)
		heap.Push(&q, int(math.Sqrt(float64(v))))
	}
	var sum int64
	for _, val := range gifts {
		sum += int64(val)
	}
	return sum
}
