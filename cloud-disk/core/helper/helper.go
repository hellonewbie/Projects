package helper

import (
	"cloud-disk/core/define"
	"crypto/md5"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
)

func MD5(s string) string {
	c := md5.New()
	c.Write([]byte(s))
	cipherStr := c.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func GenerateToken(id int, identity string, Name string) (string, error) {
	//id
	//identity
	//name
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     Name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	//密钥加密
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
