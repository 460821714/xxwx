package controllers

import (
	"xxwx/models"

	"github.com/dchest/captcha"
)

var (
	capId string
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	capId = captcha.NewLen(4)
	c.Data["CapId"] = capId
	c.Lay_tpl()

}

func (c *LoginController) Post() {
	u_name := c.Input().Get("username")
	u_pass := c.Input().Get("password")
	u_yzm := c.Input().Get("yzm")
	if captcha.VerifyString(capId, u_yzm) {
		us, _ := models.Login(u_name, u_pass)
		if us != nil {
			c.TplName = "index.html"
		} else {
			c.Data["Err"] = "用户名或密码错误，请重试"
			c.Lay_tpl()
		}
	} else {
		c.Data["Err"] = "验证码错误，请重试"
		c.Lay_tpl()
	}

}

func (c *LoginController) Lay_tpl() {
	c.Layout = "login.html"
	c.TplName = "login-captcha.html"
}
