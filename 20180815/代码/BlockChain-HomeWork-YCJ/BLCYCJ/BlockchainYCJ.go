package BLCYCJ

type BlockchainYCJ struct {
	Blocks []*BlockYCJ //存储有序的区块
}

// 增加区块到区块链里面
//结构体里面方法，增加新区块到区块链上
func (blc *BlockchainYCJ)AddBlockToBlockchainYCJ(data string,height int64,preHash []byte)  {
	//创建新区快
	newBlock := NewBlockYCJ(data,height,preHash)
	//往链中添加区块
	blc.Blocks = append(blc.Blocks,newBlock)
}

func CreateBlockchainWithGenesisBlockYCJ() *BlockchainYCJ  {
	genesisBlock := CreateGenesisBlockYCJ("Genesis Block")
	return  &BlockchainYCJ{[]*BlockYCJ{genesisBlock}}
}
