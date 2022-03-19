package main

import (
	"fmt"
	"github.com/hackerMao/go_example/retriever/mock"
	"github.com/hackerMao/go_example/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{Content: "this is a fake imooc.com"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
