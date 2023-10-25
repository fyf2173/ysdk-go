package util

import (
	"runtime"
	"testing"
)

func TestGetRandBetween(t *testing.T) {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//t.Log(r.Int63n(15))
	for _, v := range []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} {
		t.Log(v)
		t.Log(GetRandBetween(1, 7))
	}
	t.Log(runtime.Version())
}
