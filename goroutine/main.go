package main

import (
	"fmt"
	"sync"
)

func test() {
	for i := 0; i < 10000; i++ {
		go func() {
			fmt.Println("hello ", i)
		}()
	}
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			wg.Add(1)
			fmt.Println("hello ", i)
		}(i)
	}

	fmt.Println("main")
	wg.Wait()
	//time.Sleep(time.Second)
}
