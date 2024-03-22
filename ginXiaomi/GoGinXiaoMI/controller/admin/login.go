package admin

import (
	"fmt"
	"ginXiaomi/GoGinXiaoMI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})

}
func (con LoginController) DoLogin(c *gin.Context) {

	captchaId := c.PostForm("captchaId")

	verifyValue := c.PostForm("verifyValue")

	if flag := models.VerifyCaptcha(captchaId, verifyValue); flag == true {
		c.String(http.StatusOK, "验证码验证成功")
	} else {
		c.String(http.StatusOK, "验证码验证失败")
	}

}

func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, err := models.MakeCaptcha()

	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}
