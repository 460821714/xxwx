package config

import "github.com/astaxie/beego"

var (
	Mysql_url string //数据库链接
)

func init() {
	Mysql_url = beego.AppConfig.String("Mysql_url")

}
