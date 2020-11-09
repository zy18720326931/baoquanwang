package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
)

type KycController struct {
	beego.Controller
}

func (k *KycController)Get(){
	k.TplName="kyc.html"




}
func (k *KycController)Post(){

var user models.User
	err:=k.ParseForm(&user)
	if err != nil {
		k.Ctx.WriteString("用户数据出取失败")
		return
	}
	_,err=user.UpdataUser()
	if err != nil {
		k.Ctx.WriteString("用户实名认证失败")
		return
	}
	Files, err := models.QueryRecordByPhone(user.Phone)
	if err != nil {

		k.Ctx.WriteString("身份验证出错")
		return
	}

	k.Data["Files"] = Files
	k.Data["Phone"]=user.Phone
   k.TplName="new.html"
}