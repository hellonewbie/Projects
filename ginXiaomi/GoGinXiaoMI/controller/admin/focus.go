package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FocusController struct {
	BaseController
}

func (con FocusController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{})
}
func (con FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}
func (con FocusController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{})
}

func (con FocusController) Delete(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/login.html", gin.H{})
}
