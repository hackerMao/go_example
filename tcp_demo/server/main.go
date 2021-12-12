package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/**
tcp server
*/

func process(conn net.Conn) {
	var buf [256]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn fieled, error: %v\n", err)
			return
		}
		fmt.Println(string(buf[:n]))
		fmt.Printf("回复：")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("write data fieled, error:%v\n", err)
			return
		}
	}
}

func main() {
	// 绑定端口并监听
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("start server on 0.0.0.0:8080 fieled, error: %v\n", err)
		return
	}
	// 接收连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accepted fieled, error: %v\n", err)
			return
		}
		// 与客户端通信
		go process(conn)
	}
}
