package util

import (
	"fmt"
	"math/rand"
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

func TestMathRand(t *testing.T) {
	var randResult = make(map[int64]int)
	for i := 0; i <= 105; i++ {
		r := rand.Int63n(100)
		fmt.Println("-->", r)
		randResult[r]++
	}
	t.Log(randResult)
}
