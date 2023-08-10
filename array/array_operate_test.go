package array

import (
	"fmt"
	"testing"
)

type Arr1 struct {
	Name string
}

func TestArrayKeys(t *testing.T) {
	var arr = []Arr1{{Name: "a"}, {Name: "b"}}
	t.Logf("%#v", GetColumns(arr, "Name"))
}

func TestArrayPluckByT(t *testing.T) {
	var tmpArr1 = []string{"a", "b", "c", "d", "e", "f", "g"}
	t.Logf("tmpArr1 pluck %+v", ArrayPluckByT(tmpArr1, 3))

	var tmpArr2 = []int{1, 2, 3, 4, 5, 6, 7, 8}
	t.Logf("tmpArr2 pluck %+v", ArrayPluckByT(tmpArr2, 3))

	var tmpArr3 []Arr1
	for i := 0; i <= 8; i++ {
		tmpArr3 = append(tmpArr3, Arr1{Name: fmt.Sprintf("%d", i)})
	}
	t.Logf("tmpArr3 pluck %v", ArrayPluckByT(tmpArr3, 3))
}
