package p0238

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		Input    []int
		Expected []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{24, 12, 8, 6},
		},
		{
			[]int{-1, 1, 0, -3, 3},
			[]int{0, 0, 9, 0, 0},
		},
	}

	for _, test := range tests {
		got := productExceptSelf(test.Input)
		if !reflect.DeepEqual(got, test.Expected) {
			t.Fatalf(
				"Input: %v, Expected: %v, Got: %v",
				test.Input,
				test.Expected,
				got)
		}
	}
}
