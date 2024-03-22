package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (c BaseController) Success(ctx *gin.Context) {
	ctx.String(http.StatusOK, "成功")
}
func (c BaseController) Error(ctx *gin.Context) {
	ctx.String(http.StatusOK, "失败")
}
