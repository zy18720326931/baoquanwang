package controllers

import "github.com/astaxie/beego"

type ResControllers struct {
	beego.Controller
}

func (w *ResControllers) Get() {
	w.TplName = "login.html"
}
