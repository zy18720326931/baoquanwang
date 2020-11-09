package controllers

import (
	"DataCertProject/models"
	"DataCertProject/nuli"
	"github.com/astaxie/beego"
	"time"
)

type SendesmsController struct {
	beego.Controller
}

func (s *SendesmsController) Post() {
	/*
		1,获取前端数据
	*/
	var Smsrecord models.Sendsms
	if err := s.ParseForm(&Smsrecord); err != nil {
		s.Ctx.WriteString("验证码输入失败")
		return
	}
	code := nuli.GenValidateCode(6)
	result, err := nuli.SendSms(Smsrecord.Phone, code, nuli.SMS_TPL_LOGIN)

	if err != nil {
		s.Ctx.WriteString("验证码发送失败")
		return
	}
	//调用成功，发送失败
	if result.Code != "OK" {
		s.Ctx.WriteString(result.Message)
		return
	}
	Smsrecord.Code= code
	Smsrecord.BizId = result.BizId
	Smsrecord.Message = result.Message
	Smsrecord.Status = result.Code
	Smsrecord.TimeStamp = time.Now().Unix()
	_, err = Smsrecord.SetSmstomysql()
	if err != nil {
		s.Ctx.WriteString("验证数据存储失败")
		return
	}

	  s.Data["BizId"] =result.BizId
	  s.Data["Phone"]=Smsrecord.Phone
	  s.TplName= "login_sms_submit.html"

}
