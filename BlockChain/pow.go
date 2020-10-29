package BlockChain

import (
	"DataCertProject/nuli"
	"bytes"
	"crypto/sha256"
	"math/big"
)

//比特币挖矿难度值: 两周
const DIFICULTYFF= 16

/**
 * 工作量证明结构体
 */
type ProofOfWork struct {
	//目标值
	Target *big.Int
	//工作量证明算法对应的哪个区块
	Block Block
}

/**
 * 实例化一个pow算法实例
 */
func NewPoW(block Block) ProofOfWork {
	target := big.NewInt(1)             //初始值
	target.Lsh(target, 255-DIFICULTYFF) //左移
	pow := ProofOfWork{
		Target: target,
		Block:  block,
	}
	return pow
}

/*
 * pow算法：寻找符合条件的nonce值
 */

func (p ProofOfWork) Run() (int64) {
	var nonce int64
	//var bigBlock *big.Int //声明
	bigBlock := new(big.Int)
	for {

		block := p.Block

		heightBytes, _ := nuli.Inttobyte(block.Height)
		timeBytes, _ :=  nuli.Inttobyte(block.TimeStamp)
		versionBytes := nuli.Stringtobyte(block.Version)
		nonceBytes, _ :=  nuli.Inttobyte(nonce)
		//bytes.Join函数，用于[]byte的拼接
		blockBytes := bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.Data,
			block.PrevHash,
			versionBytes,
			nonceBytes,
		}, []byte{})
		sha256Hash := sha256.New()
		sha256Hash.Write(blockBytes)
		block256Hash := sha256Hash.Sum(nil)

		//sha256hash(区块+nonce) 对应的大整数
		bigBlock = bigBlock.SetBytes(block256Hash)

		if p.Target.Cmp(bigBlock) == 1 { //如果满足条件时，退出循环
			break
		}
		nonce++ //如果条件不满足，nonce值+1，继续下次循环

	}
	return nonce
}
