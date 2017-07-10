package matchers

import (
	"encoding/xml"
	"errors"
	"net/http"
	"fmt"
	"log"
	"regexp"
	"../search"
)

/**
	搜索rss源的匹配器
 */

type (
	item struct {
		XMLName xml.Name `xml:"item"`
		PubDate string `xml:"pubDate"`
		Title string `xml:"title"`
		Description string `xml:"description"`
		Link string `xml:"link"`
		GHUID string `xml:"ghuid"`
		GeoRssPoint string `xml:"georss:point"`
	}

	image struct {
		XMLName xml.Name `xml:"image"`
		URL string `xml:"url"`
		Title string `xml:"title"`
		Link string `xml:"link"`
	}

	channel struct {
		XMLName xml.Name `xml:"channel"`
		Title string `xml:"title"`
		Description string `xml:"description"`
		Link string `xml:"link"`
		PubDate string `xml:"pubDate"`
		LastBuildDate string `xml:"lastBuildDate"`
		TTL string `xml:"ttl"`
		Language string `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

type rssMatcher struct {

}

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// retrive发送HTTP Get请求获取rss数据源并解码
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return  nil, errors.New("No rss feed URI provided")
	}

	//获取rss数据源
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}
	// 一旦从函数返回,关闭返回的响应链接
	defer resp.Body.Close()

	// 检查状态码是否为200
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Respone Error %d\n", resp.StatusCode)
	}

	// 将rss数据源解码到我们定义的结构类型
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}

func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error)  {
	var results []*search.Result
	log.Printf("Search Feed Type[%s] site[%s] for URI[%s]\n", feed.Type, feed.Name, feed.URI)
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}
	for _, channelItem := range document.Channel.Item {
		// 检查标题部分是否包含搜索项
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return  nil, err
		}
		if matched {
			results = append(results, &search.Result{
				Field: "Title",
				Content:channelItem.Title,
			})
		}
		// 检查描述部分是否包含搜索项
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}
		if matched {
			results = append(results, &search.Result{
				Field:"Description",
				Content: channelItem.Description,
			})
		}
	}
	return results, nil
}
