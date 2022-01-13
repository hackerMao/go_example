package main

import (
	"flag"
	"fmt"
)

func main() {
	var host string
	flag.StringVar(&host, "host", "0.0.0.0", "127.0.0.1")
	var port int
	flag.IntVar(&port, "port", 8080, "-p 8000")
	flag.Parse()

	fmt.Printf("address: %s:%d\n", host, port)
}
