package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiRoutesInit(router *gin.Engine) {
	apiRoute := router.Group("/api")
	{
		apiRoute.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"username": "张三", "age": 20})
		})
		apiRoute.GET("/news", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"title": "这是新闻"})
		})
	}
}
