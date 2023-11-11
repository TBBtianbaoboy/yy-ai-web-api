//---------------------------------
//File Name    : to_string.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-09 22:16:04
//Description  :
//----------------------------------
package utils

import "strconv"

//Int64ToString
func Int64ToString(value int64) string {
	return strconv.FormatInt(value,10)
}

//IntToString
func IntToString(value int) string {
	return strconv.Itoa(value)
}
