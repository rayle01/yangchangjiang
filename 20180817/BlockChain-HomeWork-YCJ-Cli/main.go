package main
import (
	"./BLCYCJ"
)
func main() {
	////创世区块，存db测试
	//blockchain := BLCYCJ.CreateBlockChainWithGenesisBlockYCJ("第二个区块吧")
	//blockchain.PrintYCJChains()

	//命令行添加创世区块测试
	blockchain:=BLCYCJ.CreateBlockChainWithGenesisBlockYCJ("我去，要成功了")
	cli:=BLCYCJ.CliYCJ{blockchain}
	cli.RunYCJ()

}