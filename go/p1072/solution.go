package p1072

func maxEqualRowsAfterFlips(matrix [][]int) int {
	var n int

	patToN := map[string]int{}
	for _, col := range matrix {
		p := make([]byte, len(col))
		if col[0] != 0 {
			for i, w := range col {
				p[i] = '0' + byte(w)
			}
		} else {
			for i, w := range col {
				p[i] = '0' + byte(w^1)
			}
		}
		k := string(p)
		patToN[k]++
		if m := patToN[k]; m > n {
			n = m
		}
	}

	return n
}
