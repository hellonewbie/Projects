package picture

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"smallProGarm/CtroMySQL"
	"smallProGarm/ManageToken"
)

type Picture struct {
	Picureurl string `form:"pictureurl"`
}

func GetPicture(r *gin.Engine, DB *gorm.DB, ReClient *redis.Client) {
	r.POST("/getpicture", func(ctx *gin.Context) {
		token := ctx.PostForm("token")
		fmt.Println("--------------" + token + "______________")
		openid, err := ManageToken.ParseGetToken(ReClient, token)
		if err != nil {
			log.Print("GetPicture ParseGetToken failed")
		}
		ctx.JSON(http.StatusOK, gin.H{
			"AllPicture": CtroMySQL.GetIdPicture(DB, openid),
			"token":      token,
		})
	})

}
