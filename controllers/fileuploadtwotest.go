package controllers

import (
	"DataCertProject/BlockChain"
	"DataCertProject/models"
	"DataCertProject/nuli"
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"time"
)

type TwoController struct {
	beego.Controller
}

func (t *TwoController) Post() {
	//获取标题
	name := t.Ctx.Request.PostFormValue("hears")
	Phone := t.Ctx.Request.PostFormValue("phone")
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
	defer thefile.Close()
	//计算文件哈希
	filena, err := os.Open(uploaddir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer filena.Close()
	write := bufio.NewWriter(thefile)
	//new一个write
	filesize, err := io.Copy(write, f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("文件大小:", filesize)

	filestr := nuli.Md5hashfile(filena)
	//	//fmt.Println("文件哈希:", filestr)
	//t.Ctx.WriteString("上传成功~！！！！！！！")

	Tofile := models.UploadRecord{}
	Tofile.FileName = h.Filename
	Tofile.FileCert = filestr
	Tofile.FileSize = filesize
	Tofile.Phone = Phone
	Tofile.CertTime = time.Now().Unix()
	Tofile.FileTitle = name
	_, err = models.UploadRecord.SaveRecord(Tofile)
	if err != nil {
		fmt.Println(err.Error())
		t.Ctx.WriteString("数据传入出错")
		return
	}

	Files, err := models.QueryRecordByPhone(Phone)
	if err != nil {
		fmt.Println(err.Error())
		t.Ctx.WriteString("身份验证出错")
		return
	}
	  Users,err:=models.QueryUserByPhone(Phone)
	if err != nil {
		fmt.Println(err.Error())
		t.Ctx.WriteString("用户认证信息传入错误")
		return
	}
    Sha256forfile:=nuli.SHA256Hashforread(filena)
	Userdata := models.Corddata{
    Baoquanid:[]byte(filestr),
    Crethash: []byte(Sha256forfile),
    Username:Users.Username,
    Phone:Users.Phone,
    CordId:Users.Cardid,
    Filename:h.Filename,
    Filesize:filesize,
    CreTime:time.Now().Unix(),

	}
   UserdataByte,err:= Userdata.NewencordforCorddata()
	if err != nil {
		fmt.Println(err.Error())
		t.Ctx.WriteString("用户认证信息序列化出现错误")
		return
	}


	//这一步是将data放入文件里，而这里不应该放入只有一个文件。。。。
	_, err = BlockChain.CHAIN.Severblock(UserdataByte)
	if err != nil {
		fmt.Println(err.Error())
		t.Ctx.WriteString("区块未存入")
		return
	}
	fmt.Println(Files)
	fmt.Println(Tofile)

	t.Data["Files"] = Files
	t.Data["Phone"] = Phone
	t.TplName = "new.html"

	//解决问题后调到上面,把数据传到页面

	//下一步
}
