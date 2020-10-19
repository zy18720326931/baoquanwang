package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user_register", &controllers.RegisterControllers{})
	beego.Router("/login2", &controllers.RequControllers{})
	beego.Router("/login3", &controllers.ResControllers{})
	beego.Router("/Upload", &controllers.TwoController{})
	//下面是跳转界面转至login2 （ 新增界面）
	beego.Router("/Tologin2",&controllers. TonewController{})
}
