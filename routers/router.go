package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterControllers{})
	beego.Router("/user_register", &controllers.RegisterControllers{})
	beego.Router("/login2", &controllers.RequControllers{})
	//直接登录下面
	beego.Router("/login3", &controllers.ResControllers{})
	beego.Router("/Upload", &controllers.TwoController{})
	//下面是跳转界面转至login2 （ 新增界面）
	beego.Router("/Tologin2",&controllers. TonewController{})
	//以下是证书页面
	beego.Router("/lookcord",&controllers.CordController{})
	///user_kyc,用于实名认证的路由
	beego.Router("/user_kyc",&controllers.KycController{})

	///user_kyc.
	beego.Router("/user_kyc.html",&controllers.KycController{})
	//login_sms.html前往短信验证码登入页面
	beego.Router("login_sms.html",&controllers.SmsController{})
	//处理数据并发送验证码
	beego.Router("/send_sms",&controllers.SendesmsController{})
	///login_sms处理得到的数据
	beego.Router("/login_sms",&controllers.SmsController{})
}
