package array

import (
	"testing"
)

type Arr1 struct {
	Name string
}

func TestArrayKeys(t *testing.T) {
	var arr = []Arr1{{Name: "a"}, {Name: "b"}}
	t.Logf("%#v", GetColumns(arr, "Name"))
}
