package main

import (
	"./BLCYCJ"
)

func main() {
	////创世区块
	//blockchain := BLCYCJ.CreateBlockChainWithGenesisBlockYCJ(nil)
	//fmt.Print(blockchain)
	cli:=BLCYCJ.CLIYCJ{}
	cli.RunYCJ()
}