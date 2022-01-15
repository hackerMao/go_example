package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var s = "hello golang"

	//编码
	encodeToString := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(encodeToString)

	//解码
	decodeBytes, _ := base64.StdEncoding.DecodeString(encodeToString)
	fmt.Println(string(decodeBytes))

	// 对URL编码
	var url = "https://www.baidu.com?keyword=golang"
	urlEncode := base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println(urlEncode)

	// 对URL解码
	urlDecodeByte, _ := base64.URLEncoding.DecodeString(urlEncode)
	fmt.Println(string(urlDecodeByte))
}
