package main

import (
	"log"
	"os"
	"io/ioutil"
	"io"
)

var (
	Trace *log.Logger
	Info *log.Logger
	Warning *log.Logger
	Error *log.Logger
)

func init()  {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("failed to open log file:", err)
	}
	Trace = log.New(ioutil.Discard, "Trace: ", log.Ldate|log.Ltime|log.Llongfile)
	Info = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Llongfile)
	Warning = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Llongfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "Error: ", log.Ldate|log.Ltime|log.Llongfile)
}

func main()  {
	Trace.Println("message")
	Info.Println("message")
	Warning.Println("message")
	Error.Println("message")
}
