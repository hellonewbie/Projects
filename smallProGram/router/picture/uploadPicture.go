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
	"smallProGarm/qiniu/GoSdk/CreateUpToken"
)

func Upload(r *gin.Engine, DB *gorm.DB, ReClient *redis.Client) {
	//HTML调试
	//r.LoadHTMLGlob("./qiniu/Gosdk/template/*")
	//r.GET("/index", func(ctx *gin.Context) {
	//	ctx.HTML(http.StatusOK, "post.html", nil)
	//})
	r.POST("/upload", func(ctx *gin.Context) {
		//从请求中读取文件
		//var token model.Token
		//err := ctx.BindJSON(&token)
		//if err != nil {
		//	log.Print("BindJson failed ")
		//	log.Fatal(err)
		//}
		token := ctx.PostForm("token")
		fmt.Println(token)
		//f, err := ctx.FormFile("f") //从请求中获取携带的参数一样
		fns := ctx.Request.MultipartForm.File["f"]

		for _, fheader := range fns {
			dst := fmt.Sprintf("./qiniu/picture/%s", fheader.Filename)
			err := ctx.SaveUploadedFile(fheader, dst)
			if err != nil {
				log.Print("SaveUploadFile failed")
			}
			code, url := CreateUpToken.UpToQiNiu(fheader)
			openid, err2 := ManageToken.ParseGetToken(ReClient, token)
			if err2 != nil {
				log.Print("ManageToken ParseGetToken failed:43")
			}
			fmt.Println(openid + "______________")
			CtroMySQL.BindIdAndPic(DB, openid, fheader.Filename, url)
			if err != nil {
				log.Print("ParseToken failed")
				log.Fatal(err)
			}
			urls := make([]string, 0)
			urls = append(urls, url)
			ctx.JSON(http.StatusOK, gin.H{
				"code":  code,
				"msg":   "ok",
				"url":   urls,
				"token": token,
			})
		}
		//if err != nil {
		//	ctx.JSON(http.StatusBadRequest, gin.H{
		//		"err": err.Error(),
		//	})

		//} else {
		//将读取的文件保存到本地（服务器本地）
		//dst := fmt.Sprintf("./qiniu/picture/%s", f.Filename)
		//fmt.Println(dst)
		//err = ctx.SaveUploadedFile(f, dst)
		//if err != nil {
		//	log.Print("SaveUploadFile failed")
		//}
		//code, url := CreateUpToken.UpToQiNiu(f)
		//fmt.Println(token)
		//openid := ManageToken.ParseGetToken(ReClient, token)
		//fmt.Println(openid + "______________")
		//CtroMySQL.BindIdAndPic(DB, openid, url)
		//if err != nil {
		//	log.Print("ParseToken failed")
		//	log.Fatal(err)
		//}
		//ctx.JSON(http.StatusOK, gin.H{
		//	"code":  code,
		//	"msg":   "ok",
		//	"url":   url,
		//	"token": token,
		//})
		//}
	})
}
