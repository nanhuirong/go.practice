package search

import (
	"os"
	"encoding/json"
)

/**
	用于读取json文件
 */

const dataFile  = "/home/neil/work/ideaProject/go.practice/src/ch2/data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI string `json:"link"`
	Type string `json:"type"`
}


func RetrieveFeeds() ([]*Feed, error){
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	// 当函数返回时关闭文件
	defer file.Close()
	// 将文件解码到切片中,切片的每一项是一个指向Feed类型的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}


