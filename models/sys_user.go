package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id           int `orm:"auto"`
	Name, Passwd string
	Total_amount float64
	Permission   int
	Creator      string
	Fxsm         string
	Phone        string
}

func Login(u string, p string) (user *User, err error) {
	sql := "SELECT id FROM user WHERE name = ? and passwd =?"
	err = orm.NewOrm().Raw(sql, u, p).QueryRow(&user)
	return user, err
}
