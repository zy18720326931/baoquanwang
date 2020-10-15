package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"path"
)

/*
*func (this *Index) Upload() {
	f, h, _ := this.GetFile("file") //获取上传的文件
	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool= map[string]bool{
		".jpg":true,
		".jpeg":true,
		".png":true,
		".bat":true,
	}
	if _,ok:=AllowExtMap[ext];!ok{
		this.Ctx.WriteString( "后缀名不符合上传要求" )
		return
	}

	//创建目录
	uploadDir := "static/upload/"
	err := os.MkdirAll( uploadDir , 777)
	if err != nil {
		this.Ctx.WriteString(fmt.Sprintf("%v",err) )
		return
	}

	//构造文件名称
	fpath := uploadDir + h.Filename
	defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况

err = this.SaveToFile("file", fpath)
	if err != nil {
this.Ctx.WriteString(fmt.Sprintf("%v",err))
	}

	this.Ctx.WriteString("上传成功~！！！！！！！")
}

*/
type FileUploadsController struct {
	beego.Controller
}

func (this *FileUploadsController) Post() {
	filename := this.Ctx.Request.PostFormValue("hear")
	fmt.Println(filename)
	f, h, _ := this.GetFile("myfile") //获取上传的文件
	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".bat":  true,
	}
	if _, ok := AllowExtMap[ext]; !ok {
		this.Ctx.WriteString("后缀名不符合上传要求")
		return
	}

	//创建目录
	uploadDir := "static/upload/"
	err := os.MkdirAll(uploadDir, 777)
	if err != nil {
		this.Ctx.WriteString(fmt.Sprintf("%v", err))
		return
	}
	hashMd5 := md5.New()
	hashMd5.Write([]byte(h.Filename)) //获得结构体user中的用户密码并粉碎
	bytes := hashMd5.Sum(nil)
	h.Filename = hex.EncodeToString(bytes)
	//构造文件名称
	fpath := uploadDir + h.Filename

	defer f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况

	err = this.SaveToFile("myfile", fpath)
	if err != nil {
		this.Ctx.WriteString(fmt.Sprintf("%v", err))
		return
	}

	this.Ctx.WriteString("上传成功~！！！！！！！")
	this.TplName = "new.html"
}
