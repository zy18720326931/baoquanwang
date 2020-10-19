package models

import (
	"DataCertProject/db"
	"fmt"
)

type Filedata struct {
	Id          int
	File_name   string
	File_size   int64
	File_string string
	File_title  string
	Cre_time    int64
	Phone       string
}

//file,head,err:=getfile("name“)
//io.reader类型的file指针？
//os.openfile(路径，执行动作，权限)
//io.copy(里面是writer类型，reader类型)都是接口
//返回文件大小和err
//这时候文件已经放入文件夹copy到
//文件hash
//

func (u Filedata) SavetoMysql() (int64, error) {
	rs, err := db.Db.Exec("insert into user_files(file_name,file_size,file_string,file_title,cre_time,Phone)"+"values(?,?,?,?,?,?)",
		u.File_name, u.File_size, u.File_string, u.File_title, u.Cre_time, u.Phone)
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	id, err := rs.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	return id, nil

}
func Request(Phone string) ([]Filedata, error) {
	rows, err := db.Db.Query("select id,file_name,file_size,file_string,file_title,cre_time,Phone from user_files where Phone=?", Phone)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var Files Filedata
	Filesdata := make([]Filedata, 0)
	for rows.Next() {
		err = rows.Scan(&Files.Id, &Files.File_name, &Files.File_size, &Files.File_string, &Files.File_title, &Files.Cre_time, &Files.Phone)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		Filesdata = append(Filesdata, Files)
	}

	return Filesdata, nil
}
