package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("hello worker2")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func worker1(ctx context.Context) {
	go worker2(ctx)
LOOP:
	for {
		fmt.Println("hello worker1")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	ctx, cancle := context.WithCancel(context.Background())
	wg.Add(1)
	go worker1(ctx)
	time.Sleep(time.Second * 3)
	cancle()
	wg.Wait()
	fmt.Println("over")
}
