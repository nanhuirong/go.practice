package main

import (
	"time"
	"log"
	"../runner"
	"os"
)

const timeout  = 10 * time.Second

func main()  {
	log.Println("starting work.")
	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("interrupt")
			os.Exit(2)
		}
	}
	log.Println("end")
}

func createTask() func(int)  {
	return func(id int) {
		log.Printf("task %d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
