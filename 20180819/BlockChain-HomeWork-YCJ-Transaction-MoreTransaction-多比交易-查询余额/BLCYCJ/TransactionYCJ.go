package BLCYCJ

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
	"encoding/hex"
)

//定义交易的数据
type TransactionYCJ struct {
	//1.交易ID-->就是交易的Hash
	TxIDYCJ []byte
	//2.输入
	VinsYCJ []*TxInputYCJ
	//3.输出
	VoutsYCJ []*TxOutputYCJ
}

/*
交易：
1.CoinBase交易：创世区块中
2.转账产生的普通交易：
 */

 func NewCoinBaseTransactionYCJ(address string) *TransactionYCJ{
 	txInput:=&TxInputYCJ{[]byte{},-1,"Genesis Data"}
 	txOutput:=&TxOutputYCJ{10,address}
 	txCoinBaseTransaction:=&TransactionYCJ{[]byte{},[]*TxInputYCJ{txInput},[]*TxOutputYCJ{txOutput}}
 	//设置交易ID
 	txCoinBaseTransaction.SetIDYCJ()
 	return txCoinBaseTransaction
 }

 //交易ID--->根据tx，生成一个hash
 func (tx *TransactionYCJ) SetIDYCJ(){
 	//1.tx--->[]byte
 	var buf bytes.Buffer
 	encoder:=gob.NewEncoder(&buf)
 	err:=encoder.Encode(tx)
 	if err != nil{
 		log.Panic(err)
	}
 	//2.[]byte-->hash
 	hash:=sha256.Sum256(buf.Bytes())
 	//3.为tx设置ID
 	tx.TxIDYCJ = hash[:]
 }


//根据转账的信息，创建一个普通的交易
func NewSimpleTransaction(from, to string, amount int64,bc *BlockchainYCJ,txs []*TransactionYCJ) *TransactionYCJ {
	//1.定义Input和Output的数组
	var txInputs []*TxInputYCJ
	var txOuputs [] *TxOutputYCJ

	//2.创建Input
	/*
	创世区块中交易ID：c16d3ad93450cd532dcd7ef53d8f396e46b2e59aa853ad44c284314c7b9db1b4
	 */

	//获取本次转账要使用output
	total,spentableUTXO := bc.FindSpentableUTXOsYCJ(from,amount,txs) //map[txID]-->[]int{index}

	for txID,indexArray:=range spentableUTXO{
		txIDBytes,_:=hex.DecodeString(txID)
		for _,index:=range indexArray{
			txInput := &TxInputYCJ{txIDBytes, index, from}
			txInputs = append(txInputs, txInput)
		}
	}


	//idBytes, _ := hex.DecodeString("c16d3ad93450cd532dcd7ef53d8f396e46b2e59aa853ad44c284314c7b9db1b4")
	//idBytes, _ := hex.DecodeString("143d7db0d5cce24645edb2ba0b503fe15969ade0c721edfd3578cd731c563a16")
	//txInput := &TxInput{idBytes, 1, from}
	//txInputs = append(txInputs, txInput)

	//3.创建Output

	//转账
	txOutput := &TxOutputYCJ{amount, to}
	txOuputs = append(txOuputs, txOutput)

	//找零
	txOutput2 := &TxOutputYCJ{total - amount, from}
	txOuputs = append(txOuputs, txOutput2)

	//创建交易
	tx := &TransactionYCJ{[]byte{}, txInputs, txOuputs}

	//设置交易的ID
	tx.SetIDYCJ()
	return tx

}

//判断tx是否时CoinBase交易
func (tx *TransactionYCJ) IsCoinBaseTransactionYCJ() bool {

	return len(tx.VinsYCJ[0].TxIDYCJ) == 0 && tx.VinsYCJ[0].VoutYCJ == -1
}
