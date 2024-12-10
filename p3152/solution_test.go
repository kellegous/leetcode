package p3152

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		Nums     []int
		Queries  [][]int
		Expected []bool
	}{
		{
			[]int{3, 4, 1, 2, 6},
			[][]int{{0, 4}},
			[]bool{false},
		},
		{
			[]int{4, 3, 1, 6},
			[][]int{{0, 2}, {2, 3}},
			[]bool{false, true},
		},
		{
			[]int{1},
			[][]int{{0, 0}},
			[]bool{true},
		},
	}

	for _, test := range tests {
		if got := isArraySpecial(test.Nums, test.Queries); !reflect.DeepEqual(test.Expected, got) {
			t.Fatalf("Test w/ nums = %#v and queries = %#v failed: expected %#v but got %#v",
				test.Nums,
				test.Queries,
				test.Expected,
				got)
		}
	}
}
