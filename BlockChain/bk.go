package BlockChain

import (
	"DataCertProject/models"
	"errors"
	"fmt"
	"github.com/boltdb/bolt-master"
)

const TONG = "thebuck"
const LAST_KEY = "last_buket"

var CHAIN  BlockChain
/*
**区块链结构体定意义：用于代表一条链
*1,将新产生的区块于已有的连起来
*可以查询区块信息
*可以遍历所有区块
 */
type BlockChain struct {
	LastBlock []byte
	Boltdb    *bolt.DB
}

func Newblockchain() BlockChain {

	var blc BlockChain
	bb, err := bolt.Open("where.db", 0600, nil)
	if err != nil {
		panic("请查看连接！！！")
	}
	//获得创世区块

	//	//如果存在呢？
	bb.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(TONG))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(TONG))
			if err != nil {
				panic(err.Error())
			}
		}

		lastvalue := bucket.Get([]byte(LAST_KEY))

		if lastvalue == nil {
			block := CretegansisBlock()
			blockbyte, _ := block.Newencoder()
			err = bucket.Put(block.Hash, blockbyte)
			if err != nil {
				panic("创世区块存储失败，请重试!!!!")
			}
			err = bucket.Put([]byte(LAST_KEY), block.Hash)
			if err != nil {
				panic("新区块更新失败，请重试!!!!")
			}
			blc = BlockChain{
				LastBlock: block.Hash,
				Boltdb:    bb,
			}

		} else {

			blockbyte := bucket.Get(lastvalue)
			lastblock, err := Newdecoder(blockbyte)
			if err != nil {

				panic("新区块hash获取失败，请重试!!!!")
			}
			blc = BlockChain{
				LastBlock: lastblock.Hash,
				Boltdb:    bb,
			}

		}

		return nil

	})
	CHAIN=blc
	return blc
}

// 该方法可以将新生成的区块放入桶中
//由于上面的改良下方函数都可以的到新的区块
func (bthis BlockChain) Severblock(data []byte) (Block, error) {
	var e error
	boltdb := bthis.Boltdb
	blocks := new(Block)
	boltdb.View(func(tx *bolt.Tx) error {
		theb := tx.Bucket([]byte(TONG))

		if theb == nil {
			e = errors.New("boltdb未创建，请重试!")
			return e
		}
		blockbyte := theb.Get(bthis.LastBlock)

		//反序列化
		blocks, _ = Newdecoder(blockbyte)

		return nil
	})
	fmt.Println("最新区块的Hash值:%x\n", blocks.Hash)

	newblock := NewBlock(blocks.Height+1, data, blocks.Hash)

	boltdb.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(TONG))
		if bk == nil { //将数据写入同种
			e = errors.New("桶未创建！！")
			return e

		}
		blockbyte, _ := newblock.Newencoder()
		bk.Put(newblock.Hash, blockbyte)

		bk.Put([]byte(LAST_KEY), newblock.Hash)
		bthis.LastBlock = newblock.Hash
		return nil
	})
	return newblock, e
}

func (bc BlockChain) Qureyblock(height int64) *Block {
	//通过区块高度查询区块信息
	if height < 0 {
		return nil
	}
	var block *Block
	db := bc.Boltdb
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(TONG))
		if bucket == nil {
			panic("查询失败")
		}
		hashkey := bc.LastBlock
		for {
			lastbytes := bucket.Get(hashkey)
			eachbk, _ := Newdecoder(lastbytes)
			if eachbk.Height < height {
				return nil
			}
			if eachbk.Height == height {
				block = eachbk
				break
			}
			hashkey = eachbk.PrevHash
		}
		return nil
	})

	return block
}
func (bc BlockChain) Qureallblock()([]*Block,error) {
	var e error
	allblocks := make([]*Block, 0)
	blockbolt := bc.Boltdb
	blockbolt.View(func(tx *bolt.Tx) error {
		thebucket := tx.Bucket([]byte(TONG))
		if thebucket == nil {
			e = errors.New("没有找到这个桶")
			return e
		}
		lasthash := thebucket.Get([]byte(LAST_KEY))

		for {

			lastblockbytes := thebucket.Get([]byte(lasthash))
			eachblock, _ := Newdecoder(lastblockbytes)
			allblocks = append(allblocks, eachblock)
			if eachblock.Height == 0 {
				break
			}
            lasthash=eachblock.PrevHash
		}

		return nil
	})
return allblocks,e
}
func (bc BlockChain)QureForid(cretid []byte)*Block  {
	block :=new(Block)
	blockbolt := bc.Boltdb

	blockbolt.View(func(tx *bolt.Tx) error {
		thebucket := tx.Bucket([]byte(TONG))
		if thebucket == nil {

			return nil
		}
		lasthash := thebucket.Get([]byte(LAST_KEY))
         var BYtes  *models.Corddata
		for  {
			lastblockbytes := thebucket.Get([]byte(lasthash))
			eachblock, _ := Newdecoder(lastblockbytes)
             BYtes,_= models.NewdecordforCorddata(eachblock.Data)
			if string(BYtes.Baoquanid)==string(cretid) {
				block=eachblock
				break
			}
			lasthash=eachblock.PrevHash
			
		}

		return nil
	})


	return  block
}