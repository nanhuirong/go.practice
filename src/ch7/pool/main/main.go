package main

import (
	"log"
	"io"
	"sync/atomic"
	"sync"
	"../../pool"
	"time"
	"math/rand"
)

const (
	maxGoroutines = 25
	pooledResources = 2
)

type dbConnecttion struct {
	ID int32
}

func (dbConn *dbConnecttion) Close() error  {
	log.Println("Close: connection", dbConn.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("create: new connection", id)
	return &dbConnecttion{id}, nil
}

func main()  {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}
	for query := 0; query < maxGoroutines ; query++ {
		go func(q int) {
			perfromQueries(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()
	log.Println("shutdown")
	p.Close()
}

func perfromQueries(query int, pool *pool.Pool)  {
	conn, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer pool.Release(conn)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("queryID:%d, connID:%d\n", query, conn.(*dbConnecttion).ID)
}