package crypto

import (
	"encoding/json"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashSalt(password []byte) string {
	b, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return ""
	}
	return string(b)
}

func VerifyPassword(hashPassword, password []byte) (bool, error) {
	var err error
	if err = bcrypt.CompareHashAndPassword(hashPassword, password); err == nil {
		return true, nil
	}
	return false, err
}

func GenerateToken(username string, realIp string) string {
	var data = make(map[string]interface{})
	data["ip"] = realIp
	data["t"] = time.Now().Unix()
	data["username"] = username
	b, _ := json.Marshal(data)
	return Md5Str(string(b))
}
