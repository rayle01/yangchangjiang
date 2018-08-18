package BLCYCJ
import (
	"bytes"
	"encoding/binary"
	"log"
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