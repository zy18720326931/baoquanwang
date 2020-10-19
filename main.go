package main

import (
	"DataCertProject/db"
	_ "DataCertProject/routers"
	"github.com/astaxie/beego"
)

func main()  {

	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	db.ConDB()
	beego.Run()



}