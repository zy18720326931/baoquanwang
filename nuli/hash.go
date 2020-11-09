package nuli

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

func Md5hashstring(data string) string {
	hashMd5 := md5.New()
	hashMd5.Write([]byte(data)) //获得结构体user中的用户密码并粉碎
	bytes := hashMd5.Sum(nil)
	return hex.EncodeToString(bytes)
}

func Md5hashfile(reader io.Reader) string {
	filebyte, _ := ioutil.ReadAll(reader)

	hashMd5 := md5.New()
	hashMd5.Write([]byte(filebyte))
	bytes := hashMd5.Sum(nil)
	return hex.EncodeToString(bytes)
}
func SHA256Hash(data []byte) ([]byte) {
	//、对数据进行sha256
	sha256Hash := sha256.New()
	sha256Hash.Write(data)
	return sha256Hash.Sum(nil)
}
func SHA256Hashforread(reader io.Reader)string {
	filebyte, _ := ioutil.ReadAll(reader)
	sha256Hash := sha256.New()
	sha256Hash.Write(filebyte)
	return hex.EncodeToString(sha256Hash.Sum(nil))


}