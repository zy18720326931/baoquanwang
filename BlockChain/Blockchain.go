package BlockChain

import (
	"DataCertProject/nuli"
	"encoding/gob"
	"time"

	"bytes"
)

/**
 *  区块结构体的定义
 */
type Block struct {
	Height    int64  //区块高度
	TimeStamp int64  //时间戳
	Hash      []byte //区块的hash
	Data      []byte // 数据
	PrevHash  []byte //上一个区块的Hash
	Version   string //版本号
	Nonce     int64  //随机数，用于pow工作量证明算法计算
}

func CretegansisBlock( ) Block {
	block:=NewBlock(0, []byte{},[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})

return block

}
/**
 * 新建一个区块实例，并返回该区块
 */
func NewBlock(height int64, data []byte, prevHash []byte) Block {
	//1、构建一个block实例，用于生成区块
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
		Version:   "0x01",
	}

	//2.为新生成的block，寻找合适的nonce值
	zhouchengkai := NewPoW(block)
	nonce := zhouchengkai.Run()

	//3.将block的Nonce设置为找到的合适的nonce数
	block.Nonce = nonce

	//调用util.SHA256Hash进行hash计算
	/**
	 * 问题分析：
		① util.SHA256Hash要求一个[]byte参数
		② block是一个自定义结构体, 与①类型不匹配
	 * 解决思路：将block结构体转换为[]byte类型数据
	 * 方案：
		① block结构体中包含6个字段，其中3个已经是[]byte
	    ② 只需将剩余3个字段转换为[]byte类型
	    ③ 将6个字段[]byte进行拼接即可
	*/
	heightBytes, _ := nuli.Inttobyte(block.Height)
	timeBytes, _ := nuli.Inttobyte(block.TimeStamp)
	versionBytes := nuli.Stringtobyte(block.Version)
	nonceBytes, _ := nuli.Inttobyte(block.Nonce)
	//bytes.Join函数，用于[]byte的拼接
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		timeBytes,
		data,
		prevHash,
		versionBytes,
		nonceBytes,
	}, []byte{})

	//4、设置第7个字段hash
	block.Hash = nuli.SHA256Hash(blockBytes)

	return block
}
func (block Block)Newencoder()( []byte,error){
	buff:=new(bytes.Buffer)
    err:=gob.NewEncoder(buff).Encode(block)
	if err !=nil {
		return nil,err
	}
	return  buff.Bytes() ,nil
}
func Newdecoder(b []byte)(*Block,error){
    var block *Block
	err:=gob.NewDecoder(bytes.NewReader(b)).Decode(&block)
	if err !=nil {
		return nil,err
	}
	return block,nil

}