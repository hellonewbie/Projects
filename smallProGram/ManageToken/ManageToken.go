package ManageToken

import (
	"fmt"
	"github.com/go-redis/redis"
	"smallProGarm/jwt"
	"smallProGarm/model"
	"smallProGarm/utils"
)

func MagToken(Client *redis.Client, Resp *model.WXLoginResp) string {
	JwtParam := new(jwt.JwtParam)
	//设置Redis连接
	JwtClient := JwtParam.SetRedisCache(Client)
	SetTokenKey := JwtParam.SetDefaultSecretKey("Eternal") //resp.SessionKey
	err := JwtParam.JwtInit(JwtClient, SetTokenKey)
	if err != nil {
		utils.Mylog.Error("content", "jwt init failed", "error", err.Error())
	}
	token, _ := jwt.CreateToken(Resp.OpenId) //resp.OpenId
	fmt.Println(token)
	fmt.Println(jwt.ParseToken(token))
	//jwt.UnsetToken(token)
	return token
}

//解析拿到的token
func ParseGetToken(ReClient *redis.Client, token string) (string, error) {
	JwtParam := new(jwt.JwtParam)
	//设置Redis连接
	JwtClient := JwtParam.SetRedisCache(ReClient)
	SetTokenKey := JwtParam.SetDefaultSecretKey("Eternal")
	err := JwtParam.JwtInit(JwtClient, SetTokenKey)
	if err != nil {
		utils.Mylog.Error("content", "jwt init failed", "error", err.Error())
	}
	openid, err2 := jwt.ParseToken(token)
	if err2 != nil {
		fmt.Println("ParseToken failed")
	}
	fmt.Println(openid + "------------------------")
	return openid, err2
}
