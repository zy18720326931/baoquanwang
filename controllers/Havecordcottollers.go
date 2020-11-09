package controllers

import (
	"DataCertProject/BlockChain"
	"DataCertProject/models"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type CordController struct {
	beego.Controller
}

func (C *CordController) Get() {
	certid := C.GetString("Certid")
	theblock := BlockChain.CHAIN.QureForid([]byte(certid))

	Certdata, err := models.NewdecordforCorddata(theblock.Data)
	Certdata.Crethashstr=string(Certdata.Crethash)
	Certdata.Baoquanidstr=strings.ToUpper(string(Certdata.Baoquanid))
	if err != nil {
		C.Ctx.WriteString("数据读取失败")
		return
	}




	fmt.Println(Certdata)
	C.Data["Certdata"] = Certdata
	C.TplName = "Forhavecrod.html"
}
