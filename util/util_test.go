package util

import (
	"fmt"
	"testing"
)

func TestTrimFloatZeroSuffix(t *testing.T) {
	t.Log(fmt.Sprintf("%.3f", float64(10)/100))
	t.Log(TrimFloatZeroSuffix(fmt.Sprintf("%.2f", float64(10)/100)))
}

func TestCompareVersion(t *testing.T) {
	t.Log(CompareVersion2("1.01", "1.001"))
}
