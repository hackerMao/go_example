package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file fieled, err:%v\n", err)
		return
	}
	// 获取文件类型
	fmt.Printf("%T\n", file)
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info fieled, err:%v\n", err)
		return
	}
	fmt.Printf("文件大小：%d B\n", fileInfo.Size())
}
