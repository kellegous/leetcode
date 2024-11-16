package p1574

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		Input    []int
		Expected int
	}{
		{
			Input:    []int{1, 2, 3, 10, 4, 2, 3, 5},
			Expected: 3,
		},
		{
			Input:    []int{5, 4, 3, 2, 1},
			Expected: 4,
		},
		{
			Input:    []int{1, 2, 3},
			Expected: 0,
		},
		{
			Input:    []int{1, 3, 2, 4},
			Expected: 1,
		},
	}

	for _, test := range tests {
		if got := findLengthOfShortestSubarray(test.Input); got != test.Expected {
			t.Fatalf("expected: %v, got: %v (arr = %v)", test.Expected, got, test.Input)
		}
	}
}
