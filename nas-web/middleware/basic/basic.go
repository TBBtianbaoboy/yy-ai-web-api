//---------------------------------
//File Name    : basic.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-18 14:48:24
//Description  :
//----------------------------------
package basic

import "strings"

// CheckURL 判定路径是否存在于配置文件的忽略列表中
func CheckURL(path string, urlList []string) bool {
	for _, item := range urlList {
		if path == item || strings.Contains(path, "static") {
			return true
		}
	}
	return false
}
