package p0010

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		S        string
		P        string
		Expected bool
	}{
		{
			S:        "aa",
			P:        "a",
			Expected: false,
		},
		{
			S:        "aa",
			P:        "a*",
			Expected: true,
		},
		{
			S:        "aab",
			P:        "c*a*b",
			Expected: true,
		},
		{
			S:        "mississippi",
			P:        "mis*is*p*.",
			Expected: false,
		},
		{
			S:        "aaa",
			P:        "a*a",
			Expected: true,
		},
	}

	for _, test := range tests {
		got := isMatch(test.S, test.P)
		if got != test.Expected {
			t.Fatalf(
				"s = \"%v\", p = \"%v\" expected %t got %t",
				test.S,
				test.P,
				test.Expected,
				got)
		}
	}
}
