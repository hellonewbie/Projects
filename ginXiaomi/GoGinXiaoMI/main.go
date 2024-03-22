package main

import (
	"ginXiaomi/GoGinXiaoMI/models"
	"ginXiaomi/GoGinXiaoMI/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/**/*")
	routes.AdminRoutesInit(r)
	models.InitMysql()
	var DB models.RedisClient
	DB.RedisInit()
	r.Run(":8080")
}
