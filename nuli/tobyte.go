package nuli

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func Inttobyte(theint int64)([]byte,error){
	buff:=new(bytes.Buffer)
	err:=binary.Write(buff,binary.BigEndian,theint)
	if err !=nil {
		fmt.Println(err.Error())
		return nil,err
	}
	return buff.Bytes(),nil
}

func Stringtobyte(thestring string) []byte {
	return []byte(thestring)
}