package main

import (
	"runtime"
	"sync"
	"fmt"
	"sync/atomic"
)

var (
	counter int64
	wg sync.WaitGroup
)

func incCounter(id int)  {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		value := counter
		// 当前goroutine从线程退出,放回队列
		runtime.Gosched()
		value++
		counter = value
	}
}

func incCounterAtomic(id int)  {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}

func main()  {
	// 分配一个逻辑处理器给调度者
	//runtime.GOMAXPROCS(1
	// 给每一个可用的核心分配一个逻辑处理器
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//runtime.GOMAXPROCS(2)
	//var wg sync.WaitGroup
	//wg.Add(2)
	//
	//fmt.Println("starting goroutines")
	//
	//go func() {
	//	defer wg.Done()
	//	for count := 0; count < 3;count++  {
	//		for char := 'a'; char < 'a' + 26; char++ {
	//			fmt.Printf("%c ", char)
	//		}
	//	}
	//}()
	//go func() {
	//	defer wg.Done()
	//	for count := 0; count < 3;count++  {
	//		for char := 'A'; char < 'A' + 26; char++ {
	//			fmt.Printf("%c ", char)
	//		}
	//	}
	//}()
	//fmt.Println("waiting to finish")
	//wg.Wait()
	//fmt.Println("\nend")
	//fmt.Println(runtime.NumCPU())

	// 包含竞争状态
	wg.Add(2)
	//go incCounter(1)
	//go incCounter(2)
	go incCounterAtomic(1)
	go incCounterAtomic(2)
	wg.Wait()
	fmt.Println(counter)
}
