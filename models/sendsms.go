package models

import "DataCertProject/db"

type Sendsms struct {
	BizId     string `form:"biz_id"`    //业务号
	Phone     string `form:"phone"`     //手机号
	Code      string `form:"code"`      //验证码
	Status    string `form:"status"`    //阿里云状态码
	Message   string `form:"message"`   //短信sdk调用描述信息
	TimeStamp int64  `form:"timestamp"` //时间戳
}

func (s *Sendsms) SetSmstomysql() (int64, error) {
	rt, err := db.Db.Exec("insert into user_sms(bizid,phone,code,status,message,timestamps) values(?,?,?,?,?,?)",
		s.BizId, s.Phone, s.Code, s.Status, s.Message, s.TimeStamp)
	if err != nil {
		return -1, err
	}
	return rt.RowsAffected()

}

func (s *Sendsms) Querybybizid() (*Sendsms, error) {
	var sendsmsrecode Sendsms
	rw := db.Db.QueryRow("select bizid,phone,code,status,message,timestamps from user_sms where phone=? and bizid=? and code=?", s.Phone, s.BizId, s.Code)
	err := rw.Scan(&sendsmsrecode.BizId, &sendsmsrecode.Phone, &sendsmsrecode.Code, &sendsmsrecode.Status, &sendsmsrecode.Message, &sendsmsrecode.TimeStamp)
	if err != nil {
		return nil, err
	}
	return &sendsmsrecode, nil
}
