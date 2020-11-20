package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
)

var base58Alphabets  =[]byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
func main() {
	var arr  =[]byte("何日旭大变态！")
	cipherText:=Base58Encode(arr)
	fmt.Println("base58编码后：",string(cipherText))
	fmt.Println("base58解码后：",string(Base58Decode(cipherText)))
}

func Base58Encode(input []byte)[]byte  {
	var result  []byte
	x:=big.NewInt(0).SetBytes(input)
	base:=big.NewInt(int64(len(base58Alphabets)))
	zero:=big.NewInt(0)
	mod:=&big.Int{}
	for x.Cmp(zero)!=0  {
		x.DivMod(x,base,mod)
		result=append(result,base58Alphabets[mod.Int64()])
	}
	if input[0]==0x00 {
		result=append(result,base58Alphabets[0])
	}
	ReverseBytes(result)
	return result
}

func Base58Decode(input []byte)[]byte  {
	result:=big.NewInt(0)
	for  _,b:=range input  {
		charIndex:=bytes.IndexByte(base58Alphabets,b)
		result.Mul(result,big.NewInt(58))
		result.Add(result,big.NewInt(int64(charIndex)))
	}
	decoded:=result.Bytes()
	if input[0]==base58Alphabets[0] {
		decoded=append([]byte{0x00},decoded...)
	}
	return decoded
}

func Base58EncodeHexString(input string) string {
	arr,_:=hex.DecodeString(input)
	res:=Base58Encode(arr)
	return fmt.Sprintf("%s",res)
}



func ReverseBytes(data []byte)  {
	for i,j:=0,len(data)-1 ;i <j;i,j=i+1,j-1 {
		data[i],data[j]=data[j],data[i]
	}
}