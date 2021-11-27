package p0010

type state byte

const (
	unknown  state = 0
	knownYes state = 1
	knownNo  state = 2
)

func boolToState(v bool) state {
	if v {
		return knownYes
	}
	return knownNo
}

func (s state) isKnown() bool {
	return s != unknown
}

func match(i int, j int, s string, p string, mt []state) bool {
	ix := i*(len(p)+1) + j
	if mt[ix].isKnown() {
		return mt[ix] == knownYes
	}

	if j == len(p) {
		ans := i == len(s)
		mt[ix] = boolToState(ans)
		return ans
	}

	// does character match?
	m := i < len(s) && (p[j] == s[i] || p[j] == '.')

	var ans bool
	// is this repeating?
	if j+1 < len(p) && p[j+1] == '*' {
		ans = match(i, j+2, s, p, mt) || (m && match(i+1, j, s, p, mt))
	} else {
		ans = m && match(i+1, j+1, s, p, mt)
	}

	mt[ix] = boolToState(ans)
	return ans
}

func isMatch(s string, p string) bool {
	return match(0, 0, s, p, make([]state, (len(s)+1)*(len(p)+1)))
}
