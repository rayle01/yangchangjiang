package BLCYCJ

//定义TxOutput的结构体
type TxOutputYCJ struct {
	//金额
	ValueYCJ int64  //金额
	//锁定脚本，也叫输出脚本，公钥，目前先理解为用户名，钥花费这笔前，必须钥先解锁脚本
	ScriptPubKeyYCJ string
}

//判断TxOutput是否时指定的用户解锁
func (txOutput *TxOutputYCJ) UnlockWithAddress(address string) bool{
	return txOutput.ScriptPubKeyYCJ == address
}
