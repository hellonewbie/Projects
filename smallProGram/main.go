package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"smallProGarm/initMatters"
	"smallProGarm/router/login"
	"smallProGarm/router/picture"
)

func main() {
	r := gin.Default()
	//获取MySQL、Redis连接
	DB, ReClient := initMatters.InitMatters()
	//登入的路由返回code解析后的信息
	login.LoginRouter(r, ReClient)
	//文件上传路由绑定文件和openid
	picture.Upload(r, DB, ReClient)
	//获取可以管理的照片
	picture.GetPicture(r, DB, ReClient)
	//删除图片
	picture.DeletePicture(r, DB, ReClient)
	//测试
	err := r.Run(":8080")
	if err != nil {
		log.Print("Run failed")
		log.Fatal(err)
	}
}
