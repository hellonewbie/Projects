package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(s string) string {
	c := md5.New()
	c.Write([]byte(s))
	cipherStr := c.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
