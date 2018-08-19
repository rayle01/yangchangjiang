package BLCYCJ

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
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
