package controllers

import (
	"DataCertProject1/nuli"
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
)

type TwoController struct {
	beego.Controller
}

func (t *TwoController) Post() {
	//获取标题
	name := t.Ctx.Request.PostFormValue("hear")
	fmt.Println(name)
	//

	f, h, err := t.GetFile("myfile")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	uploaddir := "static/upload/" + h.Filename
	thefile, err := os.OpenFile(uploaddir, os.O_RDWR|os.O_CREATE, 777)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	write := bufio.NewWriter(thefile)
	//new一个write
	filesize, err := io.Copy(write, f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("文件大小:", filesize)

	filestr := nuli.Md5hashfile(f)
	fmt.Println("文件哈希:", filestr)
	t.Ctx.WriteString("上传成功~！！！！！！！")
	t.TplName = "new.html"
}
