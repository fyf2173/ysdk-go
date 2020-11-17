package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	yapi "github.com/fyf2173/ysdk-go/api"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func HashSalt(password []byte) string {
	b, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return ""
	}
	return string(b)
}

func VerifyPassword(hashPassword, password []byte) bool {
	var err error
	if err = bcrypt.CompareHashAndPassword(hashPassword, password); err == nil {
		return true
	}
	fmt.Println("auth failed, error: ", err.Error())
	return false
}

func GenerateToken(username string) string {
	var data = make(map[string]interface{})
	data["ip"] = yapi.GetRealIp()
	data["t"] = time.Now().Unix()
	data["username"] = username
	b, _ := json.Marshal(data)
	hash := md5.New()
	hash.Write(b)
	return hex.EncodeToString(hash.Sum(nil))
}
