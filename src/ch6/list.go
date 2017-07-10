package main

import (
	"sync"
	"fmt"
	"time"
	"sync/atomic"
)

var (
	shutdown int64
	wg sync.WaitGroup
)

func doWork(name string)  {
	defer wg.Done()

	for  {
		fmt.Printf("doing %s work\n", name)
		time.Sleep(250 * time.Microsecond)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("shutting %s down\n", name)
			break
		}
	}
}

func main()  {
	wg.Add(2)
	go doWork("A")
	go doWork("B")
	time.Sleep(1 * time.Second)
	fmt.Println("shutdown")

	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()

}
