package main

import (
	"DataCertProject/BlockChain"
	"DataCertProject/db"
	_ "DataCertProject/routers"
	"github.com/astaxie/beego"
)

//const BUCKET_NAME = "class3"

func main() {


	BlockChain.Newblockchain()
	//1.连接数据库
	db.ConDB()
	//2.静态资源路径设置
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	beego.Run()
}
