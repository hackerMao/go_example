package main

import "fmt"

func funcA() {
	fmt.Println("连接数据库")
}

/**
recover一定要放在panic之前，且搭配defer使用，否则捕获不到错误
*/
func funcB() {
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()
	panic("出现错误")
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
