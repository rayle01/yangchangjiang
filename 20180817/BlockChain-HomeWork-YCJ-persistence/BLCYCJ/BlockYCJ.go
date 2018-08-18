package BLCYCJ

import (
	"time"
	"fmt"
	"bytes"
	"log"
	"encoding/gob"
)
//区块结构体
// height,prevBlockHash
//
type BlockYCJ struct {
	//1、区块高度 64int
	HeightYCJ int64
	//2、上一个区块的HASH
	PrevBlockHashYCJ []byte
	//3、交易数据 64位数组
	DataYCJ []byte
	//4、时间戳
	TimstampYCJ int64
	//5、Hash byte数组
	HashYCJ [] byte
	//int值
	NonceYCJ int64
}


// 1、创建新区块
// @param data 交易值
// @param height 高度
// @param height 前一个区块的hash
func NewBlockYCJ(data string,prevBlockHash []byte,height int64) *BlockYCJ {
	//创建区块
	block := &BlockYCJ{height,prevBlockHash,[]byte(data),time.Now().Unix(),nil,0}
	//调用工作量证明的方法并且返回有效的Hash和Nonce值

	pow := NewProofOfWorkYCJ(block)
	//00000 工作量证明
	//两个返回值的函数，计算hash和nonce
	hash,nonce := pow.RunYCJ()
	//赋值
	block.HashYCJ=hash
	//赋值
	block.NonceYCJ=nonce

	fmt.Println()
	return block
}
// 2、单独方法，生成创世区块
func CreateGenesisBlockYCJ(data string)  *BlockYCJ{
	return NewBlockYCJ(data,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},1)
}


//定义block的方法，用于序列化该block对象，获取[]byte
func (block *BlockYCJ) SerializeYCJ()[]byte{
	//1.创建一个buff
	var buf bytes.Buffer

	//2.创建一个编码器
	encoder:=gob.NewEncoder(&buf)

	//3.编码
	err:=encoder.Encode(block)
	if err != nil{
		log.Panic(err)
	}

	return buf.Bytes()
}

//定义一个函数，用于将[]byte，转为block对象，反序列化
func DeserializeBlockYCJ(blockBytes [] byte) *BlockYCJ{
	var block BlockYCJ
	//1.先创建一个reader
	reader:=bytes.NewReader(blockBytes)
	//2.创建解码器
	decoder:=gob.NewDecoder(reader)
	//3.解码
	err:=decoder.Decode(&block)
	if err != nil{
		log.Panic(err)
	}
	return &block
}

