package routes

import "github.com/gin-gonic/gin"

func DefaultRoutesInit(router *gin.Engine) {
	defaultRoute := router.Group("/")
	{
		defaultRoute.GET("/", func(c *gin.Context) {
			c.String(200, "首页")
		})
	}
}
