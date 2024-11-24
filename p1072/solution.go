package p1072

import (
	"math"
	"math/bits"
)

type bitset struct {
	vec []uint64
	n   int
}

func newBitset(col []int) *bitset {
	s := zeros(len(col))
	for i, v := range col {
		if v != 0 {
			s.add(i)
		}
	}
	return s
}

func zeros(n int) *bitset {
	m := int(math.Ceil(float64(n) / 64.0))
	return &bitset{
		vec: make([]uint64, m),
		n:   n,
	}
}

func (s *bitset) add(i int) {
	wx := i / 64
	bx := i % 64
	s.vec[wx] |= (1 << bx)
}

func (s *bitset) contains(i int) bool {
	wx := i / 64
	bx := i % 64
	return (s.vec[wx] & (1 << bx)) != 0
}

func (s *bitset) invert() *bitset {
	c := &bitset{
		vec: make([]uint64, len(s.vec)),
		n:   s.n,
	}
	for i, w := range s.vec {
		c.vec[i] = ^w
	}
	c.clamp()
	return c
}

func (s *bitset) count() int {
	var c int
	for _, w := range s.vec {
		c += bits.OnesCount64(w)
	}
	return c
}

func (s *bitset) clamp() {
	n := len(s.vec) - 1
	s.vec[n] &= (1<<(s.n%64) - 1)
}

func setAnd(a, b *bitset) *bitset {
	c := &bitset{
		vec: make([]uint64, len(a.vec)),
		n:   a.n,
	}
	for i, n := 0, len(a.vec); i < n; i++ {
		c.vec[i] = a.vec[i] & b.vec[i]
	}
	return c
}

func setOr(a, b *bitset) *bitset {
	c := &bitset{
		vec: make([]uint64, len(a.vec)),
		n:   a.n,
	}
	for i, n := 0, len(a.vec); i < n; i++ {
		c.vec[i] = a.vec[i] | b.vec[i]
	}
	return c
}

func setXor(a, b *bitset) *bitset {
	c := &bitset{
		vec: make([]uint64, len(a.vec)),
		n:   a.n,
	}
	for i, n := 0, len(a.vec); i < n; i++ {
		c.vec[i] = a.vec[i] ^ b.vec[i]
	}
	c.clamp()
	return c
}

func and(a, b []int) []int {
	n := len(a)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		c[i] = a[i] & b[i]
	}
	return c
}

func xor(a, b []int) []int {
	n := len(a)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		c[i] = a[i] ^ b[i]
	}
	return c
}

func or(a, b []int) []int {
	n := len(a)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		c[i] = a[i] | b[i]
	}
	return c
}

func count(a []int) int {
	var c int
	for _, v := range a {
		c += v
	}
	return c
}

func invert(a []int) []int {
	n := len(a)
	c := make([]int, n)
	for i, v := range a {
		c[i] = (^v & 1)
	}
	return c
}

func maxEqualRowsAfterFlips(matrix [][]int) int {
	// n := len(matrix[0])
	// return reduce(matrix, make([]int, n), n)

	n := len(matrix[0])
	sets := make([]*bitset, len(matrix))
	for i, col := range matrix {
		sets[i] = newBitset(col)
	}
	return maxEqualRowsAfterSetFlips(sets, zeros(n), n)
}

func maxEqualRowsAfterSetFlips(
	matrix []*bitset,
	diff *bitset,
	n int,
) int {
	a := setOr(diff, setXor(matrix[0], matrix[1]))
	b := setOr(diff, setXor(matrix[0].invert(), matrix[1]))
	if len(matrix) == 2 {
		return n - min(a.count(), b.count())
	}
	return max(
		maxEqualRowsAfterSetFlips(matrix[1:], a, n),
		maxEqualRowsAfterSetFlips(matrix[1:], b, n),
	)
}

func reduce(matrix [][]int, diff []int, n int) int {
	a := or(diff, xor(matrix[0], matrix[1]))
	b := or(diff, xor(invert(matrix[0]), matrix[1]))
	if len(matrix) == 2 {
		return n - min(count(a), count(b))
	}
	return max(
		reduce(matrix[1:], a, n),
		reduce(matrix[1:], b, n),
	)
}
