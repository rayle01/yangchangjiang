package BLCYCJ

import (
	"flag"
	"os"
	"log"
	"fmt"
)

type CLIYCJ struct {
	//BlockChain *BlockChain
}

func (cli *CLIYCJ) RunYCJ() {

	/*
	Usage:
		addblock -data DATA
		printchain


	./bc printchain
		-->执行打印的功能

	 ./bc send -from '["wangergou"]' -to '["lixiaohua"]' -amount '["4"]'
	./bc send -from '["wangergou","rose"]' -to '["lixiaohua","jace"]' -amount '["4","5"]'


	 */
	isValidArgs()

	//1.创建flagset命令对象

	CreateBlockChainCmd:=flag.NewFlagSet("createblockchain",flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	//2.设置命令后的参数对象
	flagCreateBlockChainData:=CreateBlockChainCmd.String("address","GenesisBlock","创世区块的信息")

	flagSendFromData:=sendCmd.String("from","","转账源地址")
	flagSendToData:=sendCmd.String("to","","转账目标地址")
	flagSendAmountData:=sendCmd.String("amount","","转账金额")

	//3.解析
	switch os.Args[1] {
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchain":
		err := CreateBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)

	}
	//4.根据终端输入的命令执行对应的功能
	if sendCmd.Parsed() {
		//fmt.Println("添加区块。。。",*flagAddBlockData)
		if *flagSendFromData == ""||*flagSendToData==""||*flagSendAmountData==""{
			fmt.Println("转账信息有误。。")
			printUsage()
			os.Exit(1)
		}
		//添加区块
		//cli.AddBlockToBlockChain([]*Transaction{})
		//from:=*flagSendFromData
		//to:=*flagSendToData
		//amount:=*flagSendAmountData
		from:=JSONToArray(*flagSendFromData) //[]string
		to:=JSONToArray(*flagSendToData) //[]string
		amount:=JSONToArray(*flagSendAmountData)//[]string
		//fmt.Println(from)
		//fmt.Println(to)
		//fmt.Println(amount)
		cli.Send(from,to,amount)
	}

	if printChainCmd.Parsed() {
		//fmt.Println("打印区块。。。")
		//cli.BlockChain.PrintChains()
		cli.PrintChains()
	}

	//添加创世区块的创建
	if CreateBlockChainCmd.Parsed(){
		if *flagCreateBlockChainData ==""{
			printUsage()
			os.Exit(1)
		}
		cli.CreateBlockChain(*flagCreateBlockChainData)
	}

}

//判断终端输入的参数的长度
func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

//添加程序运行的说明
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreateblockchain -address DATA -- 创建创世区块")
	fmt.Println("\tsend -from From -to To -amount Amount -- 转账交易")
	fmt.Println("\tprintchain -- 打印区块")
}


func (cli *CLIYCJ) PrintChains(){
	//cli.BlockChain.PrintChains()
	bc:=GetBlockChainObjectYCJ() //bc{Tip,DB}
	if bc == nil{
		fmt.Println("没有BlockChain，无法打印任何数据。。")
		os.Exit(1)
	}
	defer bc.DBYCJ.Close()
	bc.PrintYCJChains()
}


func (cli *CLIYCJ) AddBlockToBlockChain(txs []*TransactionYCJ){
	//cli.BlockChain.AddBlockToBlockChain(data)
	bc := GetBlockChainObjectYCJ()
	if bc == nil{
		fmt.Println("没有BlockChain，无法添加新的区块。。")
		os.Exit(1)
	}
	defer bc.DBYCJ.Close()
	bc.AddBlockToBlockChainYCJ(txs)
}

func(cli *CLIYCJ) CreateBlockChain(address string ){
	//fmt.Println("创世区块。。。")
	CreateBlockChainWithGenesisBlock(address)

}

func (cli *CLIYCJ) Send(from ,to ,amount []string){
	bc :=GetBlockChainObjectYCJ()
	if bc == nil{
		fmt.Println("没有BlockChain，无法转账。。")
		os.Exit(1)
	}
	defer bc.DBYCJ.Close()
	bc.MineNewBlockYCJ(from,to,amount)
}