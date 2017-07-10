package main

import "log"

func init()  {
	log.SetPrefix("nanhuirong: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func main()  {
	// 写到标准日志记录器
	log.Println("message")

	// Fatalln 在调用Println()之后会调用会os.Exit(1)
	//log.Fatalln("fatal message")

	// 在调用Println()之后会调用panic
	log.Panicln("panic message")
}
