package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// 解决跨域问题
// Cross-Origin Resource Sharing
func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowAllOrigins: true, //允许所有的跨域
			//允许跨域请求的方法
			AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:  []string{"*"},
			ExposeHeaders: []string{"Content-Length", "Authorization", "Content-Type"},
			//指示请求是否可以包含用户凭据（如 Cookie、HTTP 认证）进行 CORS 请求。这里设置为 true。
			AllowCredentials: true, //是否为cookie请求
			//如果跨域的不是，则返回下面这个网站
			//AllowOriginFunc: func(origin string) bool {
			//	return origin == "https://github.com"
			//}
			MaxAge: 12 * time.Hour,
		})

}
