package main

import (
	"log"
	"os"
	"./search"
	_"./matchers"

)

//在main之前调用,程序中所有的init函数都会在main之前被调用
func init()  {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main()  {
	search.Run("president")
}
