package BLCYCJ

import (
	"github.com/boltdb/bolt"
	"log"
)
//定义区块链的迭代器，专门用于迭代遍历该区块链对应的数据库中block对象
type BlockChainIteratorYCJ struct {
	DBYCJ          *bolt.DB
	CurrentHashYCJ []byte
}

func (bcIterator *BlockChainIteratorYCJ) NextYCJ() *BlockYCJ {
	block := new(BlockYCJ)
	//1.根据bcIterator，操作DB对象，读取数据库
	err := bcIterator.DBYCJ.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlockBucketName))
		if b != nil {
			//根据current获取对应的区块的数据
			blockBytes := b.Get(bcIterator.CurrentHashYCJ)
			//反序列化后得到block对象
			block = DeserializeBlockYCJ(blockBytes)
			//更新currenthash
			bcIterator.CurrentHashYCJ = block.PrevBlockHashYCJ
		}
		return nil

	})
	if err != nil {
		log.Panic(err)
	}
	return block
}
