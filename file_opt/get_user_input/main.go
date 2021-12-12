package main

import (
	"bufio"
	"fmt"
	"os"
)

func useScan() {
	fmt.Println("请输入内容：")
	var input string
	fmt.Scanln(&input)
	fmt.Printf("hello, %v\n", input)
}

func useBufIO()  {
	fmt.Println("请输入内容：")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Printf("hello, %v\n", input)
}

func main() {
	//useScan()
	useBufIO()
}
