//---------------------------------
//File Name    : hash.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-10 13:49:45
//Description  :
//----------------------------------
package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// md5String 生成MD5前缀,截断 length 长度.
func MD5String(key string, length int) string {
	h := md5.New()
	h.Write([]byte(key))
	result := hex.EncodeToString(h.Sum(nil))
	if length > len(result) {
		length = len(result)
	}
	return result[:length]
}
