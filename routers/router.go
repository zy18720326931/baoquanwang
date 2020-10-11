package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user_register",&controllers.RegisterControllers{})
    beego.Router("/login2",&controllers.RequControllers{})

}
