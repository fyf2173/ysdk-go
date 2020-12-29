package string

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// n is the len of returned rand string
func GetRandString(n int) string {
	var rands = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = rands[rand.Intn(len(rands))]
	}
	return string(b)
}

// GetRandBetween 生成区间随机数
func GetRandBetween(max, min int64) int64 {
	if min > max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + max
}
