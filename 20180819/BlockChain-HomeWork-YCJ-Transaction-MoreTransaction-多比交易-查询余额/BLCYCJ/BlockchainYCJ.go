package BLCYCJ

import (
	"github.com/boltdb/bolt"
	"fmt"
	"log"
	"os"
	"time"
	"math/big"
	"encoding/hex"
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
func (bc *BlockchainYCJ) AddBlockToBlockChainYCJ(txs []*TransactionYCJ) {
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
			newBlock := NewBlockYCJ(txs, lastBlock.HashYCJ, lastBlock.HeightYCJ+1)
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

func CreateBlockChainWithGenesisBlockYCJ(address string)  {

	/*
	1.判断数据库如果存在，直接结束方法
	2.数据库不存在，创建创世区块，并存入到数据库中
	 */
	if dbExists(){
		fmt.Println("数据库已经存在，无法创建创世区块。。")
		return
	}

	//数据库不存在
	fmt.Println("数据库不存在。。")
	fmt.Println("正在创建创世区块。。。。。")
	/*
	1.创建创世区块
	2.存入到数据库中
	 */
	//创建一个txs--->CoinBase
	txCoinBase:=NewCoinBaseTransactionYCJ(address)

	genesisBlock := CreateGenesisBlockYCJ([]*TransactionYCJ{txCoinBase})
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
	//return &BlockChain{db, genesisBlock.Hash}
}


//创建一个区块链，包含创世区块
/*
1.数据库存储，创世区块已经存在，直接返回
2.数据库不存在，创建创世区块，存入到数据库中
 */
func CreateBlockChainWithGenesisBlock(address string)  {

	/*
	1.判断数据库如果存在，直接结束方法
	2.数据库不存在，创建创世区块，并存入到数据库中
	 */
	if dbExists(){
		fmt.Println("数据库已经存在，无法创建创世区块。。")
		return
	}

	//数据库不存在
	fmt.Println("数据库不存在。。")
	fmt.Println("正在创建创世区块。。。。。")
	/*
	1.创建创世区块
	2.存入到数据库中
	 */
	//创建一个txs--->CoinBase
	txCoinBase:=NewCoinBaseTransactionYCJ(address)

	genesisBlock := CreateGenesisBlockYCJ([]*TransactionYCJ{txCoinBase})
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
	//return &BlockChain{db, genesisBlock.Hash}
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
		//fmt.Printf("\t数据：%s\n", block.DataYCJ)
		fmt.Println("\t交易信息：")
		for _,tx:=range block.Txs{
			fmt.Printf("\t\t交易ID：%s\n",tx.TxIDYCJ)
		}
		fmt.Printf("\t随机数：%d\n", block.NonceYCJ)
		//fmt.Printf("\t时间：%d\n", block.TimeStamp)
		fmt.Printf("\t时间：%s\n", time.Unix(block.TimstampYCJ, 0).Format("2006-01-02 15:04:05")) // 时间戳-->time-->Format("")

		//step2：判断block的prevBlcokhash为0,表示该block是创世取块，将诶数循环
		hashInt := new(big.Int)
		hashInt.SetBytes(block.PrevBlockHashYCJ)
		if big.NewInt(0).Cmp(hashInt) == 0 {
			break
		}

	}
}

//获取blockchainiterator的对象
//获取blockchainiterator的对象
func (bc *BlockchainYCJ) IteratorYCJ() *BlockChainIteratorYCJ {
	return &BlockChainIteratorYCJ{bc.DBYCJ, bc.TipYCJ}
}


//提供一个函数，专门用于获取BlockChain对象
func GetBlockChainObjectYCJ() *BlockchainYCJ{
	/*
		1.数据库存在，读取数据库，返回blockchain即可
		2.数据库 不存储，返回nil
	 */

	if dbExists() {
		//fmt.Println("数据库已经存在。。。")
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
	}else{
		fmt.Println("数据库不存在，无法获取BlockChain对象。。。")
		return  nil
	}
}

//新增功能：通过转账，创建区块
func (bc *BlockchainYCJ) MineNewBlockYCJ(from,to,amount []string){

	fmt.Println(from)
	fmt.Println(to)
	fmt.Println(amount)
}


//提供一个功能：查询余额
func (bc *BlockchainYCJ) GetBalanceYCJ(address string,txs[] *TransactionYCJ) int64 {
	//txOutputs := bc.UnSpent(address)
	unSpentUTXOs := bc.UnSpentYCJ(address,txs)

	var total int64
	for _, utxo := range unSpentUTXOs {
		total += utxo.Output.ValueYCJ
	}
	return total

}

//设计一个方法，用于获取指定用户的所有的未花费Txoutput
/*
UTXO模型：未花费的交易输出
	Unspent Transaction TxOutput
 */
func (bc *BlockchainYCJ) UnSpentYCJ(address string, txs [] *TransactionYCJ) []*UTXOYCJ { //王二狗
	/*
	0.查询本次转账已经创建了的哪些transaction

	1.遍历数据库，获取每个block--->Txs
	2.遍历所有交易：
		Inputs，---->将数据，记录为已经花费
		Outputs,---->每个output
	 */
	//存储未花费的TxOutput
	var unSpentUTXOs [] *UTXOYCJ
	//存储已经花费的信息
	spentTxOutputMap := make(map[string][]int) // map[TxID] = []int{vout}

	//第一部分：先查询本次转账，已经产生了的Transanction
	for i := len(txs)-1;i>=0;i--{
		unSpentUTXOs = caculate(txs[i],address,spentTxOutputMap,unSpentUTXOs)
	}



	//第二部分：数据库里的Trasacntion

	it := bc.IteratorYCJ()

	for {
		//1.获取每个block
		block := it.NextYCJ()
		//2.遍历该block的txs
		//for _, tx := range block.Txs {
		//倒序遍历Transaction
		for i := len(block.Txs) - 1; i >= 0; i-- {
			unSpentUTXOs = caculate(block.Txs[i],address,spentTxOutputMap,unSpentUTXOs)
		}

		//3.判断推出
		hashInt := new(big.Int)
		hashInt.SetBytes(block.PrevBlockHashYCJ)
		if big.NewInt(0).Cmp(hashInt) == 0 {
			break
		}

	}

	return unSpentUTXOs
}


func caculate(tx *TransactionYCJ,address string, spentTxOutputMap map[string][]int,unSpentUTXOs []*UTXOYCJ) []*UTXOYCJ{
	//遍历每个tx：txID，Vins，Vouts

	//遍历所有的TxInput
	if !tx.IsCoinBaseTransactionYCJ() { //tx不是CoinBase交易，遍历TxInput
		for _, txInput := range tx.VinsYCJ {
			//txInput-->TxInput
			if txInput.UnlockWithAddress(address) {
				//txInput的解锁脚本(用户名) 如果和钥查询的余额的用户名相同，
				key := hex.EncodeToString(txInput.TxIDYCJ)
				spentTxOutputMap[key] = append(spentTxOutputMap[key], txInput.VoutYCJ)
				/*
				map[key]-->value
				map[key] -->[]int
				 */
			}
		}
	}

	//遍历所有的TxOutput
outputs:
	for index, txOutput := range tx.VoutsYCJ { //index= 0,txoutput.锁定脚本：王二狗
		if txOutput.UnlockWithAddress(address) {
			if len(spentTxOutputMap) != 0 {
				var isSpentOutput bool //false
				//遍历map
				for txID, indexArray := range spentTxOutputMap { //143d,[]int{1}
					//遍历 记录已经花费的下标的数组
					for _, i := range indexArray {
						if i == index && hex.EncodeToString(tx.TxIDYCJ) == txID {
							isSpentOutput = true //标记当前的txOutput是已经花费
							continue outputs
						}
					}
				}

				if !isSpentOutput {
					//unSpentTxOutput = append(unSpentTxOutput, txOutput)
					//根据未花费的output，创建utxo对象--->数组
					utxo := &UTXOYCJ{tx.TxIDYCJ, index, txOutput}
					unSpentUTXOs = append(unSpentUTXOs, utxo)
				}

			} else {
				//如果map长度未0,证明还没有花费记录，output无需判断
				//unSpentTxOutput = append(unSpentTxOutput, txOutput)
				utxo := &UTXOYCJ{tx.TxIDYCJ, index, txOutput}
				unSpentUTXOs = append(unSpentUTXOs, utxo)
			}
		}
	}
	return unSpentUTXOs

}



/*
提供一个方法，用于一次转账的交易中，可以使用为花费的utxo
 */
func (bc *BlockchainYCJ) FindSpentableUTXOsYCJ(from string, amount int64,txs[]*TransactionYCJ) (int64, map[string][]int) {
	/*
	1.根据from获取到的所有的utxo
	2.遍历utxos，累加余额，判断，是否如果余额，大于等于要要转账的金额，


	返回：map[txID] -->[]int{下标1，下标2} --->Output
	 */
	var total int64
	spentableMap := make(map[string][]int)
	//1.获取所有的utxo ：10
	utxos := bc.UnSpentYCJ(from,txs)
	//2.找即将使用utxo：3个utxo
	for _, utxo := range utxos {
		total += utxo.Output.ValueYCJ
		txIDstr := hex.EncodeToString(utxo.TxID)
		spentableMap[txIDstr] = append(spentableMap[txIDstr], utxo.Index)
		if total >= amount {
			break
		}
	}

	//3.
	if total < amount {
		fmt.Printf("%s,余额不足，无法转账。。\n", from)
		os.Exit(1)
	}

	return total, spentableMap

}
