package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
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

	m.Data["Phone"] = u.Phone

	m.TplName = "login2.html"
}
