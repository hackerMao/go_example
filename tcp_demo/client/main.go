package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/**
tcp client
*/

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	defer conn.Close()
	if err != nil {
		fmt.Printf("dial 127.0.0.1:8080 fieled, error:%v\n", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("请输入:")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("write data fieled, error:%v\n", err)
			return
		}

	}
	_ = conn.Close()
}
