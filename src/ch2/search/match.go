package search

import (
	"log"
	"fmt"
)

/**
	用于支持不同匹配器的接口
 */

type Result struct {
	Field string
	Content string
}

//接口数据类型,如果包含一个方法以er结尾
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil{
		log.Println(err)
		return
	}
	for _, result := range searchResults {
		results <- result
	}
}

func Display(results chan *Result)  {
	// 通道一直阻塞,知道有结果写入
	// 一旦通道被关闭,for循环终止
	for result := range results{
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}