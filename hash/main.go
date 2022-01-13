package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	// md5
	var s string = "hello world"
	bytes := md5.Sum([]byte(s))
	// 转换成十六进制
	fmt.Printf("md5: %X\n", bytes)
	// 通过Sprintf
	md5Str := fmt.Sprintf("%X\n", bytes)
	fmt.Println(md5Str)
	// 通过哈希工具转换
	fmt.Println(hex.EncodeToString(bytes[:]))

	// 当计算大文件时可以使用下面这种方式
	m := md5.New()
	m.Write([]byte("hello "))
	m.Write([]byte("world"))
	fmt.Printf("%X\n", m.Sum(nil))

	// sha1、 sha256、 sha512、 hmac用法相似
}
