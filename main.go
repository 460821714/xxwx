package main

import (
	_ "xxwx/routers" //_ 引用该包 不能引用该包其他函数

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/xxwx?charset=utf8")
	orm.RunSyncdb("default", false, true)
	beego.Run()

}
