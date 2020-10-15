package controllers

import (
	"DataCertProject1/models"
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
	u, err := user.Querys()

	if err != nil {
		m.Ctx.WriteString("您未注册")
	}

	m.Data["Phone"] = u.Phone

	m.TplName = "login2.html"
}
