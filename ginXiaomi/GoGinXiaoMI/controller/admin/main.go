package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MainController struct {
	BaseController
}

func (con MainController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/index.html", gin.H{})
}
func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
