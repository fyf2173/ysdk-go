package crypto

import (
	"crypto/sha256"
	"strings"
	"testing"
)

func TestHashHmac(t *testing.T) {
	t.Logf("%+v \n", strings.ToUpper(HashHmac(sha256.New, "aaa", "dddddd")))
}

func TestMd5Str(t *testing.T) {
	t.Logf("%+v ", Md5Str("test"))
}
