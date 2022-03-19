package main

import (
	"fmt"
	"github.com/hackerMao/go_example/retriever/mock"
	"github.com/hackerMao/go_example/retriever/real"
	"time"
)

/**
接口变量自带指针
接口变量同样使用值传递，几乎不需要使用接口的指针
*/

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

// type switch
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Content:", v.Content)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

func main() {
	var r Retriever
	r = &mock.Retriever{Content: "this is a fake imooc.com"}
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Second,
	}
	inspect(r)
	// type assertion
	if retriever, ok := r.(mock.Retriever); ok {
		fmt.Println(retriever.Content)
	} else {
		fmt.Println("not a mock retriever")
	}

	//fmt.Println(download(r))
}
