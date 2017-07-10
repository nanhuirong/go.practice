package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

const (
	numberGoroutines = 4
	taskLoad = 10
)

var wg sync.WaitGroup

func init()  {
	rand.Seed(time.Now().Unix())
}

func worker(tasks chan string, worker int)  {
	defer wg.Done()

	for  {
		task, ok := <- tasks
		if !ok {
			//意味着通道为空,并且已经被关闭
			return
		}

		fmt.Printf("worker %d starts %s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("worker %d completed %s\n", worker, task)
	}
}

func main()  {
	tasks := make(chan string, taskLoad)
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("task %d", post)
	}
	close(tasks)
	wg.Wait()
}

