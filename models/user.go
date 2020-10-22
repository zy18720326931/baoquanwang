package models

import (

	"DataCertProject/db"
	"DataCertProject/nuli"

)

type User struct {
	Id       int    `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

/**
 * 保存用户信息的方法：保存用户信息到数据库中
 */
func (u User) SaveUser() (int64, error) {
	//1、密码脱敏处理
	u.Password =nuli.Md5hashstring(u.Password)

	//2、执行数据库操作
	row, err := db.Db.Exec("insert into user(phone, password)"+
		" values(?,?) ", u.Phone, u.Password)
	if err != nil {
		return -1, err
	}
	id, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

/**
 * 查询用户信息
 */
func (u User) QueryUser()(*User,error) {
	//1、密码脱敏处理
	u.Password = nuli.Md5hashstring(u.Password)

	row := db.Db.QueryRow("select Phone from baoquan where  Phone = ? and Password = ?",
		u.Phone, u.Password)
	err := row.Scan(&u.Phone)
	if err != nil {
		return nil,err
	}
	return &u,nil
}
