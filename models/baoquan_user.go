package models

import (
	"DataCertProject1/db"
	"DataCertProject1/nuli"
	"fmt"
)

type User struct {
	Id       int    `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

func (u User) SeveUser() (int64, error) {
	fmt.Println("执行了")
	//密码拖密处理
	/*
	*hash
	*/


	u.Password = nuli.Md5hashstring(u.Password)
	//执行数据库操作
	row, err := db.Db.Exec("insert into baoquan(phone,password)"+" values(?,?)", u.Phone, u.Password)
	if err != nil {
		return -1, err
	}
	//返回受影响的行数
	num, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return num, nil
}

func (u User) Querys() (*User,error) {

	u.Password = nuli.Md5hashstring(u.Password)
	row :=db.Db.QueryRow("select phone from baoquan where phone =? and password=? ", u.Phone, u.Password)


	err:=row.Scan(&u.Phone)
	if err !=nil {
	return nil,err
}
	return  &u,nil
}

