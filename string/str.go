package string

import "strings"

// Substr returns the substr from start to length, if length smaller than 0, Substr returns the substr from start to end
func Substr(s string, start, length int) string {
	bt := []rune(s)
	if len(bt) <= 0 {
		return ""
	}
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if length <= 0 {
		return string(bt[start:])
	}
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if !strings.Contains(haystack, needle) {
		return -1
	}
	return strings.Index(haystack, needle)
}
