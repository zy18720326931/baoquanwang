package controllers

import (
	"DataCertProject1/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterControllers struct {
	beego.Controller
}

func (r *RegisterControllers) Post() {
	fmt.Println("也执行了")
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("对不起，数据解析错误")
		return
	}
	_, err1 := user.SeveUser()
	if err1 != nil {
		fmt.Println(err1)
		r.Ctx.WriteString("对不起，用户注册失败")
		return
	}
	r.TplName = "login.html"
}
