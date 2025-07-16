package base

import (
	"net/url"
)

// UrlStringRemoveQuery 删除完整url中的 query数据
// http://www.naidu.com.com/29429407085598149.jpg?sd=ius&wd=po
// 返回数据为： http://www.naidu.com.com/29429407085598149.jpg
func UrlStringRemoveQuery(urlStr string) string {
	ud, _ := url.Parse(urlStr)
	ud.RawQuery = ""
	return ud.String()
}
