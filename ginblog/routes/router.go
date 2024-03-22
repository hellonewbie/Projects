package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

// 入口文件
func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	//什么中间件都没有使用
	r := gin.New()
	//默认加入了两个中间件engine.Use(Logger(), Recovery());gin.Default()获取到的Engine 实例集成了Logger 和 Recovery 中间件
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//用户模块的接口路由
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//文章模块的接口路由
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		//分类模块的接口路由
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		//上传文件
		auth.POST("upload", v1.Upload)
	}
	routerV1 := r.Group("api/v1")
	{
		routerV1.GET("user/get", v1.GetUsers)
		routerV1.GET("user/:id", v1.GetUserInfo)
		routerV1.GET("article/get", v1.GetArticle)
		routerV1.GET("article/list/:cid", v1.GetCateArt)
		routerV1.GET("article/info/:id", v1.GetArtInfo)
		routerV1.GET("category/get", v1.GetCategory)
		routerV1.POST("user/add", v1.AddUser)
		routerV1.POST("login", v1.Login)
	}
	r.Run(utils.HttpPort)
}
