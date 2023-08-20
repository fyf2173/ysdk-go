package util

import (
	"runtime"
	"testing"
)

func TestGetRandBetween(t *testing.T) {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//t.Log(r.Int63n(15))
	t.Log(GetRandBetween(3, 15))
	t.Log(runtime.Version())
}
