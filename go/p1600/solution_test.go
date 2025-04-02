package p1600

import (
	"reflect"
	"testing"
)

func TestXxx(t *testing.T) {
	ti := Constructor("king")
	ti.Birth("king", "andy")
	ti.Birth("king", "bob")
	ti.Birth("king", "catherine")
	ti.Birth("andy", "matthew")
	ti.Birth("bob", "alex")
	ti.Birth("bob", "asha")

	expA := []string{"king", "andy", "matthew", "bob", "alex", "asha", "catherine"}
	if got := ti.GetInheritanceOrder(); !reflect.DeepEqual(got, expA) {
		t.Fatalf(
			"Expected %v but got %v",
			expA,
			got,
		)
	}

	ti.Death("bob")

	expB := []string{"king", "andy", "matthew", "alex", "asha", "catherine"}
	if got := ti.GetInheritanceOrder(); !reflect.DeepEqual(got, expB) {
		t.Fatalf(
			"Expected %v but got %v",
			expB,
			got,
		)
	}
}
