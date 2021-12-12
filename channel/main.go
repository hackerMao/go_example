package main

import (
	"fmt"
	"sync"
)

/**
channel 必须使用make初始化后才可使用
*/

var a []int
var b chan int
var wg sync.WaitGroup

func main() {
	fmt.Println(b)     // nil
	b = make(chan int) // 不带缓冲区的channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("received number: ", <-b)
	}()
	b <- 10 //没有接受者将会卡住导致死锁

	b = make(chan int, 16) //channel 必须使用make初始化后才可使用
	fmt.Println(b)         // 指针地址
	wg.Wait()
}
