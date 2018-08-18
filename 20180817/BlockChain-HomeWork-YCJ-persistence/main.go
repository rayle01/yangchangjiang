package main

import (
	"./BLCYCJ"
)

func main() {
	////创世区块
	blockchain := BLCYCJ.CreateBlockChainWithGenesisBlockYCJ("第二个区块吧")
	blockchain.PrintYCJChains()

}