package middleware

import (
	"ginblog/utils"
	"ginblog/utils/errormsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)
var code int

//其实就是token承载的东西

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//生成token

func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	setClaims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims)
	//进行签名
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errormsg.ERROR
	}
	return token, errormsg.SUCCESS
}

//验证token

func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, errormsg.SUCCESS
	} else {
		return nil, errormsg.ERROR
	}

}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization")
		if tokenHerder == "" {
			code = errormsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errormsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.SplitN(tokenHerder, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errormsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errormsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])
		if tCode == errormsg.ERROR {
			code = errormsg.ERROR_TOEKN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errormsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errormsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errormsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key.Username)
		c.Next()
	}

}
