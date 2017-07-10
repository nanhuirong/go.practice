package search

/**
	搜索数据的默认匹配器
 */
type defaultMatcher struct {

}

func init()  {
	var matcher defaultMatcher
	Register("default", matcher)
}

// (m defaultMatcher)为方法的调用者
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error)  {
	return nil, nil
}



