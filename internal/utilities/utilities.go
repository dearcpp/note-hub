package utilities

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5EncryptToString(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}
