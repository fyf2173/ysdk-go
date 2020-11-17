package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"hash"
)

func HashHmac(f func() hash.Hash, str, key string) string {
	mac := hmac.New(f, []byte(key))
	mac.Write([]byte(str))
	return hex.EncodeToString(mac.Sum(nil))
}

func Md5Str(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}
