package util

import (
	"net"
	"strings"
	"time"
)

// TrimFloatZeroSuffix 格式化浮点数，去除小数点末尾的0或.
func TrimFloatZeroSuffix(floatStr string) string {
	dotIndex := strings.Index(floatStr, ".")
	if dotIndex < 0 {
		return floatStr
	}
	breakIndex := 0
	for i := len(floatStr) - 1; i >= dotIndex; i-- {
		tmp := string(floatStr[i])
		if tmp != "0" {
			if tmp == "." {
				i--
			}
			breakIndex = i
			break
		}
	}
	return floatStr[:breakIndex+1]
}

// Assert 断言
func Assert(err error) {
	if err != nil {
		panic(err)
	}
}

func PickUnusedPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	port := l.Addr().(*net.TCPAddr).Port
	if err := l.Close(); err != nil {
		return 0, err
	}
	return port, nil
}

func NanotimeToDatetime(nanoTime int64) string {
	return time.Unix(nanoTime/1e9, 0).Format("2006-01-02 15:04:05")
}

func TimeFixDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
