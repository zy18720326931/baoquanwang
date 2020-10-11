package controllers

import (
	"github.com/astaxie/beego"
)

type RequControllers struct {
	beego.Controller
}

func (m *RequControllers) Post() {


	m.TplName="login2.html"
}