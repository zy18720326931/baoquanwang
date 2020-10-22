package main

import (
	blockchain "DataCertProject/BlockChain"
	"DataCertProject/db"
	_ "DataCertProject/routers"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/boltdb/bolt-master"
)
const TONG  = "thebuck"
func main()  {
	block := blockchain.CretegansisBlock()
	boltdb,err:=bolt.Open("my.db",0600,nil)
	if err!=nil {
		panic("在检查一次再来吧")
	}
	defer boltdb.Close()


	boltdb.View(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(TONG))
		var blok blockchain.Block
		if b!=nil {
			blockbyte:=b.Get([]byte("0"))
			fmt.Println(blockbyte)
			err:= gob.NewDecoder(bytes.NewReader(blockbyte)).Decode(&blok)
			if err!=nil {
				fmt.Println(err.Error())
			}
			fmt.Println(blok)
		}

		return nil
	})

	boltdb.Update(func(tx *bolt.Tx) error {
		bk:=tx.Bucket([]byte(TONG))
		if bk==nil {//将数据写入同种

			b,err:=tx.CreateBucket([]byte(TONG))
			if err !=nil {
				return err
			}
			fmt.Println("桶创建成功！！")
			buff:=new(bytes.Buffer)

			err= gob.NewEncoder(buff).Encode(block)
			if err !=nil {
				return err
			}
			err=  b.Put([]byte("0"),buff.Bytes())
			if err !=nil {
				return err
			}

		}
		return nil
	})


	return
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	db.ConDB()
	beego.Run()



}