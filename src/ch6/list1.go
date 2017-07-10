package main

import (
	"sync"
	"runtime"
	"fmt"
)

var (
	counter int
	wg sync.WaitGroup
	// 互斥锁
	mutex sync.Mutex
)

func incCounterMutex(id int)  {
	defer wg.Done()

	mutex.Lock()
	{
		for count := 0; count < 2; count++ {
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
	}
	mutex.Unlock()


}

func main()  {
	wg.Add(2)
	go incCounterMutex(1)
	go incCounterMutex(2)
	wg.Wait()
	fmt.Printf("%d", counter)
}
