package main

import (
	"math/rand"
	"time"
	"sync"
	"fmt"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}

var wg sync.WaitGroup

func player(name string, court chan int)  {
	defer wg.Done()
	for  {
		ball, ok := <-court
		if !ok {
			fmt.Printf("player %s win\n", name)
			return
		}

		n := rand.Intn(100)
		if n % 13 == 0 {
			fmt.Printf("player %s missed\n", name)
			close(court)
			return
		}

		fmt.Printf("player %s hit %d\n", name, ball)
		ball++
		court <- ball
	}
}
func main()  {

	/**
	 1.无缓冲的通道,要求发送goroutine和接受goroutine同时做好准备,否则会导致阻塞
	 2.缓冲通道,无法保证消息的发送和接受在同一时间完成
	 */

	// 无缓冲区的整型通道
	//unbuffered := make(chan int)
	//// 有缓冲区的字符串通道,10指定缓冲区的大小
	//buffered := make(chan string, 10)
	//
	////通过通道发送一个字符串
	//buffered <- "nanhuirong"
	//// 从缓冲区接受值
	//value := <-buffered

	court := make(chan int)
	wg.Add(2)

	go player("nan", court)
	go player("wang", court)

	court <- 1
	wg.Wait()

}
