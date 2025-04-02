package p2257

import (
	"math"
	"math/bits"
)

type bitset []uint64

func newBitset(n int) bitset {
	m := int(math.Ceil(float64(n) / 64.0))
	return bitset(make([]uint64, m))
}

func (s bitset) add(i int) {
	wx := i / 64
	bx := i % 64
	s[wx] |= (1 << bx)
}

func (s bitset) contains(i int) bool {
	wx := i / 64
	bx := i % 64
	return (s[wx] & (1 << bx)) != 0
}

func (s bitset) count() int {
	var c int
	for _, val := range s {
		c += bits.OnesCount64(val)
	}
	return c
}

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	// thoughts: we could use union find to easily keep up with the components
	// that are unioned ... but with a bitset, we're only going to have to
	// do a popcnt on 1/64 of the set size ... which should be fine.
	indexOf := func(r int, c int) int {
		return r*n + c
	}

	k := m * n

	wgs := newBitset(k)

	s := newBitset(k)

	for _, wall := range walls {
		ix := indexOf(wall[0], wall[1])
		wgs.add(ix)
		s.add(ix)
	}

	for _, guard := range guards {
		ix := indexOf(guard[0], guard[1])
		wgs.add(ix)
		s.add(ix)
	}

	for _, guard := range guards {
		for i := guard[0] + 1; i < m; i++ {
			ix := indexOf(i, guard[1])
			if wgs.contains(ix) {
				break
			}
			s.add(ix)
		}

		for i := guard[0] - 1; i >= 0; i-- {
			ix := indexOf(i, guard[1])
			if wgs.contains(ix) {
				break
			}
			s.add(ix)
		}

		for j := guard[1] + 1; j < n; j++ {
			ix := indexOf(guard[0], j)
			if wgs.contains(ix) {
				break
			}
			s.add(ix)
		}

		for j := guard[1] - 1; j >= 0; j-- {
			ix := indexOf(guard[0], j)
			if wgs.contains(ix) {
				break
			}
			s.add(ix)
		}
	}

	return k - s.count()
}
