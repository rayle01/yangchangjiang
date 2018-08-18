package BLCYCJ
import (
	"math/big"
	"bytes"
	"crypto/sha256"
	"fmt"
)

//0000 0000 0000 0000 1001 0001 0000 .... 0001
//256位Hash里面前面至少要有16个0
const targetBit  = 16

type ProofOfWorkYCJ struct {
	Block *BlockYCJ //当前要验证的区块
	//diff int64
	targetYCJ *big.Int //大数据存储
}


//合并准备数据
func (pow *ProofOfWorkYCJ) prepareDataYCJ(nonce int) []byte  {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHashYCJ,
			pow.Block.DataYCJ,
			IntToHexYCJ(pow.Block.TimstampYCJ),
			IntToHexYCJ(int64(targetBit)),
			IntToHexYCJ(int64(nonce)),
			IntToHexYCJ(int64(pow.Block.HeightYCJ)),
		},
		[]byte{},
	)
	return data
}

func (proofOfWork *ProofOfWorkYCJ) IsValidYCJ() bool{
	//1.proofOfWork.Block.Hash
	//2.proofOfWork.target
	var hashInt big.Int
	hashInt.SetBytes(proofOfWork.Block.HashYCJ)
	//比较
	if proofOfWork.targetYCJ.Cmp(&hashInt) == 1{
		return true
	}
	return false
}

func (proofOfWork *ProofOfWorkYCJ) RunYCJ() ([]byte,int64) {
	//1 、将Block的属性拼接成字节数组
	//2、生成hash
	//3、判断hash的有效性，如果满足条件、跳出循环
	nonce := 0
	var hashInt big.Int //存储新生成的hash
	var hash [32]byte
	for{
		//准备数据
		dataBytes := proofOfWork.prepareDataYCJ(nonce)
		//生成hash
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x",hash)
		//将hash存储到hashInt
		hashInt.SetBytes(hash[:])
		//可以crtl+鼠标看源码
		// Cmp compares x and y and returns:
		if proofOfWork.targetYCJ.Cmp(&hashInt) == 1{
			break
		}
		nonce = nonce+1
	}
	return hash[:],int64(nonce)
}

//创建新的工作量证明对象
func NewProofOfWorkYCJ(block *BlockYCJ) *ProofOfWorkYCJ {
	//1、创建一个初始值为1的target

	target := big.NewInt(1)
	//2、左移动256-targetBit
	target = target.Lsh(target,256 - targetBit)

	return &ProofOfWorkYCJ{block,target}
}
