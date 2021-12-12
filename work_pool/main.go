package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	value int64
}

type result struct {
	job    *job
	result int64
}

var wg sync.WaitGroup

func provider(jc chan<- *job) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		jc <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func consumer(jc <-chan *job, rc chan<- *result) {
	defer wg.Done()
	for {
		job, ok := <-jc
		if ok {
			var sum int64 = 0
			n := job.value
			for n > 0 {
				sum += n % 10
				n = n / 10
			}
			res := &result{
				job:    job,
				result: sum,
			}
			rc <- res
		}
	}
}

func main() {
	jobChan := make(chan *job, 100)
	resultChan := make(chan *result, 100)

	wg.Add(1)
	go provider(jobChan)

	num := 24
	wg.Add(num)
	for i := 1; i <= num; i++ {
		go consumer(jobChan, resultChan)
	}

	for res := range resultChan {
		fmt.Printf("value: %d sum:%d\n", res.job.value, res.result)
	}
	wg.Wait()
}
