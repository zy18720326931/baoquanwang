package controllers

import (
	"DataCertProject/models"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type SmsController struct {
	beego.Controller
}

func (s *SmsController)Get()  {
	s.TplName="login_sms.html"
}


func (s *SmsController)Post()  {
	// 获取跳转来的信息
	var setrecode models.Sendsms
	 if err:=s.ParseForm(&setrecode);err != nil{
	 	s.Ctx.WriteString("信息跳转失败")
		 return
	 }
	 //从数据库读取信息与与前端对比
     sendsmsthis,err:=setrecode.Querybybizid()
	if err!=nil{
		s.Ctx.WriteString("未查到对应验证信息")
		return
	}
	fmt.Println(sendsmsthis)
	 us,err2:= models.QueryUserByPhone(sendsmsthis.Phone)

	if err2!=nil{
		fmt.Println(err2.Error())
		s.Ctx.WriteString("对不起用户未注册")
		return
	}

	name := strings.TrimSpace(us.Username)
	Cord := strings.TrimSpace(us.Cardid)
	Sex := strings.TrimSpace(us.Sex)
	if name==""||Cord==""||Sex=="" {
		s.Data["Phone"]=us.Phone
		s.TplName="kyc1.html"
		return
	}
	Files, err := models.QueryRecordByPhone(us.Phone)
	if err != nil {
		fmt.Println(err.Error())
		s.Ctx.WriteString("身份验证出错")
		return
	}

	s.Data["Files"] = Files
	s.Data["Phone"] = us.Phone
	s.TplName = "new.html"
}