package routers

import (
	"xxwx/controllers"

	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

func init() {
	//Router		映射 URL 到 controller
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Handler("/captcha/*.png", captcha.Server(90, 30))
}
