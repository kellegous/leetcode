package p0001

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		Nums     []int
		Target   int
		Expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, test := range tests {
		if got := twoSum(test.Nums, test.Target); !reflect.DeepEqual(got, test.Expected) {
			t.Fatalf(
				"Nums: %#v, Target: %#v, Expected: %#v, Got: %#v",
				test.Nums,
				test.Target,
				test.Expected,
				got)
		}
	}
}
