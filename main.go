package main

import (
	"DataCertProject1/db"
	_ "DataCertProject1/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/js", ".static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	db.ConDB()
	beego.Run()
}
