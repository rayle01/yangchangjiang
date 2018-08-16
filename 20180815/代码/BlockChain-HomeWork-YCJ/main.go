package main

import (
	"./BLCYCJ"
	"fmt"
)

func main() {
	////创世区块
	blockchain := BLCYCJ.CreateBlockchainWithGenesisBlockYCJ()
	//
	////新区块
	blockchain.AddBlockToBlockchainYCJ("Send 100 RMB to CJ1",blockchain.Blocks[len(blockchain.Blocks)-1].HeightYCJ+1,blockchain.Blocks[len(blockchain.Blocks)-1].HashYCJ)
	//blockchain.AddBlockToBlockchain("Send 300 RMB to CJx",blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//blockchain.AddBlockToBlockchain("Send 400 RMB to CJz",blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//blockchain.AddBlockToBlockchain("Send 200 RMB to CJf",blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//fmt.Println(blockchain)
	//fmt.Println(blockchain.Blocks)
	block := BLCYCJ.NewBlockYCJ("Test",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	fmt.Printf("%d\n",block.NonceYCJ)
	fmt.Printf("%d\n",block.HashYCJ)

	proofOfWork := BLCYCJ.NewProofOfWorkYCJ(block)
	fmt.Printf("%v",proofOfWork.IsValidYCJ())

}