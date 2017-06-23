package utils

import (

	"github.com/astaxie/beego"
)

var (
	Mysql_url string //数据库连接
	RunMode   string //运行模式
)

func init() {
	Mysql_url = beego.AppConfig.String("Mysql_url")

}
