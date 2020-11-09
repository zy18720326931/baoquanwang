package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
	"strings"
)

type RequControllers struct {
	beego.Controller
}

func (m *RequControllers) Post() {

	var user models.User
	err := m.ParseForm(&user)
	if err != nil {
		m.Ctx.WriteString("您未登录")
	}
	u, err := user.QueryUser()

	if err != nil {
		m.Ctx.WriteString("您未注册")
	}
      name := strings.TrimSpace(u.Username)
	Cord := strings.TrimSpace(u.Cardid)
	Sex := strings.TrimSpace(u.Sex)
	if name==""||Cord==""||Sex=="" {
		m.Data["Phone"]=u.Phone
		m.TplName="kyc1.html"
		return
	}
	m.Data["Phone"] = u.Phone

	m.TplName = "login2.html"
}
