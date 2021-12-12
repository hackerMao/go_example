package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x = 0
	//lock sync.Mutex
	rwLock sync.RWMutex
	wg     sync.WaitGroup
)

func read() {
	defer wg.Done()
	rwLock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond * 5)
	rwLock.RUnlock()
}

func write() {
	defer wg.Done()
	rwLock.Lock()
	x += 1
	time.Sleep(time.Millisecond * 5)
	rwLock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 20; i++ {
		go write()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start.Add(time.Second)))
}
