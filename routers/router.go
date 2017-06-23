package routers

import (
	"xxwxCMS/controllers"

	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
	beego.Handler("/captcha/*.png", captcha.Server(80, 40))
}
