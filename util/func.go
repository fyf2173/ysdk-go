package util

import "strings"

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
