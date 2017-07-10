package main

import (
	"log"
	"time"
	"../../worker"
	"sync"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task()  {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main()  {
	p := worker.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))
	for index := 0; index < 100;index++  {
		for _, name:= range names {
			np := namePrinter{
				name:name,
			}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()
}
