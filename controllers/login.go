package controllers

import (
	"xxwx/models"

	"errors"
	"fmt"

	"github.com/dchest/captcha"
)

type LoginController struct {
	//重写 beego.Controller
	BaseControllers
}

const (
	x1 = iota
	x2 = iota
	x3 = iota
)

func (login *LoginController) Get() {
	err := errors.New("emit macho dwarf:elf herder corrupted")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("x1:", x1)
	fmt.Print("x2:", x2)
	fmt.Print("x3:", x3)
	//capId := captcha.NewLen(4)
	//login.Data["CapId"] = capId
	login.TplName = "index.html"
}
func (login *LoginController) Post() {
	user := login.Input().Get("username")
	passwd := login.Input().Get("password")
	User, _ := models.Login(user, passwd)
	if User != nil {
		capId := captcha.NewLen(4)
		login.Data["CapId"] = capId
		login.TplName = "index.html"
	} else {
		login.Data["Err"] = "1"
		login.TplName = "login.html"
	}

}
