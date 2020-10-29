package main

import (
	"DataCertProject/db"
	_ "DataCertProject/routers"
	"github.com/astaxie/beego"
)

func main()  {

	//newblockchain:=BlockChain.Newblockchain()
   //
   // block,err:= newblockchain.Severblock([]byte("世界你好"))
	//if err !=nil {
	//	fmt.Println(err.Error())
	//}
   //fmt.Println(block)
   // fmt.Printf("最新区块的Hash值:%x\n",block.Hash)
   //  block3:=newblockchain.Qureyblock(1)
	//if block3==nil {
	//	fmt.Println("没有o")
	//	return
	//}
   // fmt.Println(block3.Data,block3.Height)
	//return
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	db.ConDB()
	beego.Run()



}