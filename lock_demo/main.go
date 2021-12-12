package main

import (
	"fmt"
	"sync"
)

var (
	x    = 0
	wg   sync.WaitGroup
	lock sync.Mutex
)

func add() {
	for i := 0; i < 100; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	n := 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
