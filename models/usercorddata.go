package models

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Corddata struct {
	Crethash     []byte
	Crethashstr  string
	Baoquanid    []byte
	Baoquanidstr string
	Username     string
	Phone        string
	CordId       string //身份证号
	Filename     string
	Filesize     int64
	CreTime      int64
}

func (c Corddata) NewencordforCorddata() ([]byte, error) {
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(c)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return buff.Bytes(), nil
}

func NewdecordforCorddata(this []byte) (*Corddata, error) {

	var Cda *Corddata
	err := gob.NewDecoder(bytes.NewReader(this)).Decode(&Cda)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return Cda, nil
}
