package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

//des加密
func main() {
	key:="88888888"
	data:="好好学习"
	cryptdata,_:=DesEnCrypt([]byte(data),[]byte(key))
	fmt.Printf("加密后：%x\n",cryptdata)
	//JeneritePriPem()
	return
}

func DesEnCrypt(data ,key []byte) ([]byte,error)  {
	block,err:=des.NewCipher(key)
	if err!= nil{
		return nil,err
	}
	originaldata:=PKCS5Padding(data,block.BlockSize())
	blockMode:=cipher.NewCBCEncrypter(block,key)
	cipherdata:=make([]byte,len(originaldata))
	blockMode.CryptBlocks(cipherdata,originaldata)
	return cipherdata,nil



}
func PKCS5Padding(data []byte,blocksize int) []byte  {
	padding:=blocksize-len(data)%blocksize
	paddata:=bytes.Repeat([]byte{byte(padding)},padding)
	return append(data,paddata...)
}


