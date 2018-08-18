package BLCYCJ

import (
	"github.com/boltdb/bolt"
	"fmt"
	"log"
	"os"
	"time"
	"math/big"
)

type BlockchainYCJ struct {
	//Blocks []*Block
	DBYCJ  *bolt.DB //对应的数据库对象
	TipYCJ [] byte  //存储区块中最后一个块的hash值
}
const DBName  = "blockchainYCJ.db" //数据库的名字
const BlockBucketName = "blocksYCJ" //定义bucket
// 增加区块到区块链里面
//结构体里面方法，增加新区块到区块链上
/*func (blc *BlockchainYCJ)AddBlockToBlockchainYCJ(data string,height int64,preHash []byte)  {
	//创建新区快
	newBlock := NewBlockYCJ(data,height,preHash)
	//往链中添加区块
	blc.BlocksYCJ = append(blc.Blocks,newBlock)
}*/


//添加区块到区块链中
func (bc *BlockchainYCJ) AddBlockToBlockChainYCJ(data string) {
	//1.根据参数的数据，创建Block
	//newBlock := NewBlock(data, prevBlockHash, height)
	//2.将block加入blockchain
	//bc.Blocks = append(bc.Blocks, newBlock)
	/*
	1.操作bc对象，获取DB
	2.创建新的区块
	3.序列化后存入到数据库中
	 */
	err := bc.DBYCJ.Update(func(tx *bolt.Tx) error {
		//打开bucket
		b := tx.Bucket([]byte(BlockBucketName))
		if b != nil {
			//获取bc的Tip就是最新hash，从数据库中读取最后一个block：hash，height
			blockByets := b.Get(bc.TipYCJ)
			lastBlock := DeserializeBlockYCJ(blockByets) //数据库中的最后一个区块
			//创建新的区块
			newBlock := NewBlockYCJ(data, lastBlock.HashYCJ, lastBlock.HeightYCJ+1)
			//序列化后存入到数据库中
			err := b.Put(newBlock.HashYCJ, newBlock.SerializeYCJ())
			if err != nil {
				log.Panic(err)
			}

			//更新：bc的tip，以及数据库中l的值
			b.Put([]byte("l"), newBlock.HashYCJ)
			bc.TipYCJ = newBlock.HashYCJ

		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

//创建一个区块链，包含创世区块
func CreateBlockChainWithGenesisBlockYCJ(data string) *BlockchainYCJ {

	/*
	1.判断数据库如果存在，
	2.数据库不存在，创建创世区块，并存入到数据库中
	 */
	if dbExists() {
		fmt.Println("数据库已经存在。。。")
		//打开数据库
		db, err := bolt.Open(DBName, 0600, nil)
		if err != nil {
			log.Panic(err)
		}

		var blockchain *BlockchainYCJ

		err = db.View(func(tx *bolt.Tx) error {
			//打开bucket，读取l对应的最新的hash
			b := tx.Bucket([]byte(BlockBucketName))
			if b != nil {
				//读取最新hash
				hash := b.Get([]byte("l"))
				blockchain = &BlockchainYCJ{db, hash}
			}
			return nil
		})
		if err != nil {
			log.Panic(err)
		}
		return blockchain
	}

	//数据库不存在
	fmt.Println("数据库不存在。。")
	/*
	1.创建创世区块
	2.存入到数据库中
	 */
	genesisBlock := CreateGenesisBlockYCJ(data)
	db, err := bolt.Open(DBName, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		//创世区块序列化后，存入到数据库中
		b, err := tx.CreateBucketIfNotExists([]byte(BlockBucketName))
		if err != nil {
			log.Panic(err)
		}

		if b != nil {
			err = b.Put(genesisBlock.HashYCJ, genesisBlock.SerializeYCJ())
			if err != nil {
				log.Panic(err)
			}
			b.Put([]byte("l"), genesisBlock.HashYCJ)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return &BlockchainYCJ{db, genesisBlock.HashYCJ}
}
//提供一个方法，用于判断数据库是否存在
func dbExists() bool {
	if _, err := os.Stat(DBName); os.IsNotExist(err) {
		return false
	}
	return true
}

//新增方法，用于遍历数据库，打印所有的区块
func (bc *BlockchainYCJ) PrintYCJChains() {
	/*
	.bc.DB.View(),
		根据hash，获取block的数据
		反序列化
		打印输出


	 */

	//获取迭代器
	it := bc.IteratorYCJ()
	for {
		//step1：根据currenthash获取对应的区块
		block := it.NextYCJ()
		fmt.Printf("第%d个区块的信息：\n", block.HeightYCJ+1)
		fmt.Printf("\t高度：%d\n", block.HeightYCJ)
		fmt.Printf("\t上一个区块Hash：%x\n", block.PrevBlockHashYCJ)
		fmt.Printf("\t自己的Hash：%x\n", block.HashYCJ)
		fmt.Printf("\t数据：%s\n", block.DataYCJ)
		fmt.Printf("\t随机数：%d\n", block.NonceYCJ)
		//fmt.Printf("\t时间：%d\n", block.TimeStamp)
		fmt.Printf("\t时间：%s\n", time.Unix(block.TimstampYCJ, 0).Format("2006-01-02 15:04:05")) // 时间戳-->time-->Format("")

		//step2：判断block的prevBlcokhash为0,表示该block是创世取块，将诶数循环
		hashInt := new(big.Int)
		hashInt.SetBytes(block.PrevBlockHashYCJ)
		if big.NewInt(0).Cmp(hashInt) == 0 {
			/*
			x.Cmp(y)
				-1 x < y
				0 x = y
				1 x > y
			 */
			break
		}

	}
}

//获取blockchainiterator的对象
//获取blockchainiterator的对象
func (bc *BlockchainYCJ) IteratorYCJ() *BlockChainIteratorYCJ {
	return &BlockChainIteratorYCJ{bc.DBYCJ, bc.TipYCJ}
}