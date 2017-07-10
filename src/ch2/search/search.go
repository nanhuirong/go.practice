package search

import (
	"log"
	"sync"
	_"fmt"
)

/**
 执行搜索的主控逻辑
 */

// 包级别变量,小写字母开头表示不公开,大写字母开头表示公开
// 注册用于搜索的匹配起的映射
var matchers = make(map[string]Matcher)

func Run(searchTerm string)  {
	//获取需要搜索的数据源列表
	// := 声明变量并且给变量初始值
	feeds, err := RetrieveFeeds()
	//for _, feed := range feeds {
	//	fmt.Println(feed.URI)
	//}
	//fmt.Println(len(feeds))
	if err != nil {
		log.Fatal(err)
	}
	//创建一个无缓冲通道,接受匹配后的结果
	results := make(chan *Result)

	//构建一个waitGroup,处理所有数据源
	var waitGroup sync.WaitGroup
	//设置需要等待处理,每个数据源的goroutine的数量
	waitGroup.Add(len(feeds))
	//为每个数据源启动一个goroutine来查找结果,_位置对应索引
	for _, feed := range feeds{
		// 获取一个匹配器用于查找
		// println(index)
		//for index, m := range matchers {
		//	println(index, m)
		//}
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}
		// 启动一个goroutine来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}
	// 启动一个goroutine来监控是否所有工作都完成
	go func() {
		// 等待所有任务完成
		waitGroup.Wait()
		// 用关闭通道的方式,通知Display函数
		// 退出程序
		close(results)
	}()

	//启动函数,显示返回结果并且在最后一个显示完成后返回
	Display(results)

}

func Register(feedType string, matcher Matcher)  {
	if _, exists := matchers[feedType]; exists{
		log.Fatalln(feedType, "Matcher already registered")
	}
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}

