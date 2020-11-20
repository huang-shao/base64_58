package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str:="！"
	//使用base64对str进行编码
	cipherText:=Base64EncodeString(str)
	fmt.Println("base64编码后：",cipherText)
	fmt.Println("base64解码后：",Base64DecodeString(cipherText))
	return
}
func Base64EncodeString(str string) string  {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
func Base64DecodeString(str string) string  {
	result,_:=base64.StdEncoding.DecodeString(str)
	return string(result)

}