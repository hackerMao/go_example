package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFile1() {
	file, err := os.Open("./temp.txt")
	if err != nil {
		fmt.Printf("failed to open file, error: %s\n", err)
	}
	defer file.Close()
	var buf = [256]byte{}
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%s\n", err)
			return
		}
		fmt.Println(string(buf[:n]))
		if n < 128 {
			break
		}
	}
}

func readFileByBufIO() {
	file, err := os.Open("./temp.txt")
	if err != nil {
		fmt.Printf("failed to open the file, error: %s\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("read over")
			break
		}
		if err != nil {
			fmt.Printf("read file field, err:%s", err)
			break
		}
		fmt.Println(line)
	}
}

func readFileByIOUtil() {
	ret, err := ioutil.ReadFile("./temp.txt")
	if err != nil {
		fmt.Printf("read file field, err:%s", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	// readFile1()
	// readFileByBufIO()
	readFileByIOUtil()
}
