//---------------------------------
//File Name    : ../web-utils/page.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-19 17:10:43
//Description  : 
//----------------------------------
package webutils

func GetPageStart(page, pageSize int) int {
	if page >= 1 {
		page = page - 1
	}
	return page * pageSize
}
