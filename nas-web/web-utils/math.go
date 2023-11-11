//---------------------------------
//File Name    : math.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-28 13:33:06
//Description  :
//----------------------------------
package webutils

import (
	"fmt"
	"strconv"
)

//@Func 使float64保留两位小数
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

//@Func 字节的单位转换 保留两位小数
func FormatSize(size uint64) (string) {
   if size < 1024 {
      return fmt.Sprintf("%.2fB", float64(size)/float64(1))
   } else if size < (1024 * 1024) {
      return fmt.Sprintf("%.2fKB", float64(size)/float64(1024))
   } else if size < (1024 * 1024 * 1024) {
      return fmt.Sprintf("%.2fMB", float64(size)/float64(1024*1024))
   } else if size < (1024 * 1024 * 1024 * 1024) {
      return fmt.Sprintf("%.2fGB", float64(size)/float64(1024*1024*1024))
   } else if size < (1024 * 1024 * 1024 * 1024 * 1024) {
      return fmt.Sprintf("%.2fTB", float64(size)/float64(1024*1024*1024*1024))
   } else {
      return fmt.Sprintf("%.2fEB", float64(size)/float64(1024*1024*1024*1024*1024))
   }
}
