package routers

import (
	"DataCertProject1/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user_register", &controllers.RegisterControllers{})
	beego.Router("/login2", &controllers.RequControllers{})
	beego.Router("/login3", &controllers.ResControllers{})
	beego.Router("/Upload",&controllers.TwoController{})
}
