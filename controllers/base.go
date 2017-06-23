package controllers

import (
	"xxwx/models"

	"github.com/astaxie/beego"
)

type BaseControllers struct {
	beego.Controller
	user *models.User
}

func (c BaseControllers) Prepare() {
	//	u_user := c.Ctx.GetCookie("user")
	//	u_pass := c.Ctx.GetCookie("pass")

	//c.Ctx.Redirect(302, "/login")
}
