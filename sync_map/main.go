package main

import (
	"fmt"
	"strconv"
	"sync"
)

/**
go内置的map并不是并发安全的
*/
var m = make(map[string]int)
var m2 = sync.Map{}
var lock sync.Mutex

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func f1() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := "key_" + strconv.Itoa(n)
			lock.Lock()
			set(key, n)
			lock.Unlock()
			fmt.Printf("m[%s]=%d\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func f2() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := "key_" + strconv.Itoa(n)
			m2.Store(key, n)
			value, _ := m2.Load(key)
			fmt.Printf("m[%s]=%d\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	//f1()

	f2()
}
