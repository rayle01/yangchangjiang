package BLCYCJ

type TxInputYCJ struct {
	//1.交易ID：引用的TxOutput所在的交易ID
	TxIDYCJ []byte

	//2.引用的交易中的哪个txoutput,其实就是下标
	VoutYCJ int

	//3.输入脚本，也就是解锁脚本。暂时理解为用户名
	ScriptSiqYCJ string
}

//判断TxInput是否时指定的用户消费
func (txInput *TxInputYCJ) UnlockWithAddress(address string) bool{
	return txInput.ScriptSiqYCJ == address
}