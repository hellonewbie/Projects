package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
	"net/http"
	"smallProGarm/ExpCode"
	"smallProGarm/ManageToken"
	"smallProGarm/model"
)

func LoginRouter(r *gin.Engine, RedisClient *redis.Client) *model.WXLoginResp {
	var WxResp *model.WXLoginResp
	r.POST("/login", func(ctx *gin.Context) {
		var WxCode model.Code
		err := ctx.BindJSON(&WxCode)
		if err != nil {
			log.Print("BindJson code failed")
			log.Fatal(err)
		}
		Resp, err := ExpCode.WXLogin(WxCode.Code)
		WxResp = Resp
		fmt.Println(WxCode.Code)
		if err != nil {
			log.Fatal(err)
		}
		token := ManageToken.MagToken(RedisClient, Resp)
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		fmt.Println("________________", Resp.OpenId)
	})
	return WxResp
}
