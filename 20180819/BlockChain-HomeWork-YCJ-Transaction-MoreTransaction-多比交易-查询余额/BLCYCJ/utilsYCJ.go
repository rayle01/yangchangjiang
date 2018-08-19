package BLCYCJ
import (
	"bytes"
	"encoding/binary"
	"log"
	"encoding/json"
)

//将int64转化为字节数组s
func IntToHexYCJ(num int64) []byte  {
	buff := new(bytes.Buffer)
	err := binary.Write(buff,binary.BigEndian,num)

	if err != nil{
		log.Panic(err)
	}
	return buff.Bytes()
}

/*
json解析的的函数

*/
func JSONToArray(jsonString string) []string {
	var arr [] string
	err := json.Unmarshal([]byte(jsonString), &arr)
	if err != nil {
		log.Panic(err)
	}
	return arr
}
