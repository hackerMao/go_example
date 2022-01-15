package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// 获取当前目录
	path, _ := filepath.Abs(".")
	fmt.Println(path)

	//获取当前文件的全路径
	fileFullPath, _ := filepath.Abs(os.Args[0])
	fmt.Println(fileFullPath)
	fmt.Println(exec.LookPath(os.Args[0]))

	// 获取文件最后一个元素
	fmt.Println(filepath.Base(fileFullPath))

	// 获取文件的目录
	fmt.Println(filepath.Dir(fileFullPath))

	// 获取文件扩展名
	fmt.Println(filepath.Ext("../copy_file/test.txt"))

	// 查找目录下以xx结尾的文件
	fmt.Println(filepath.Glob("../copy_file/[ab]*.txt"))

	// 遍历指定文件目录
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path, info.Name())
		return nil
	})
}
