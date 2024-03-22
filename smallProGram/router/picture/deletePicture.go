package picture

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"smallProGarm/CtroMySQL"
	"smallProGarm/ManageToken"
)

func DeletePicture(r *gin.Engine, DB *gorm.DB, ReClient *redis.Client) {
	r.DELETE("/deletePicture", func(ctx *gin.Context) {
		pictureurl := ctx.PostForm("url")
		token := ctx.PostForm("token")
		openid, err := ManageToken.ParseGetToken(ReClient, token)
		if err != nil {
			log.Print("Delete Picture ParseToken failed")
		}

		CtroMySQL.DelPicture(DB, openid, pictureurl)
		ctx.JSON(http.StatusOK, gin.H{
			"token":       token,
			"picturename": pictureurl,
		})
	})
}
