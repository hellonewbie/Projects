package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (c UserController) Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "这是用户首页")
}
func (c UserController) Add(ctx *gin.Context) {
	ctx.String(http.StatusOK, "增加用户")
}
