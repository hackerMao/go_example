package main

import (
	"fmt"
	"sync"
)

/**
channel练习：
	1、启动一个goroutine生成100个数发送ch1
	2、启动一个goroutine，从ch1中取值，计算其平方放到ch2中
	3、在main中 从ch2中取值并打印到控制台
*/

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	wg := sync.WaitGroup{}

	go product(ch1, &wg)
	go consumer(ch1, ch2, &wg)
	wg.Wait()
	for {
		v, ok := <-ch2
		if !ok {
			break
		}
		fmt.Println(v)
	}
}

func consumer(ch1 <-chan int, ch2 chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	for {
		v, ok := <-ch1

		if !ok {
			break
		}
		ch2 <- v * v
	}
	close(ch2) // 不关闭channel的话就会在读取的时候进入死循环，从而导致死锁
}

func product(ch1 chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}
