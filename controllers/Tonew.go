package controllers

import "github.com/astaxie/beego"

type TonewController struct {
	beego.Controller
}

func (t *TonewController)Get() {
	Phone := t.GetString("phone")
    t.Data["Phone"]	= Phone
    t.TplName="login2.html"
}