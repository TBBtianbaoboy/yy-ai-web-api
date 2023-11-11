//---------------------------------
//File Name    : env.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-08 11:39:54
//Description  : 
//----------------------------------
package utils

import (
	"os"
)

// GetEnvDefault
func GetEnvDefault(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return value
}

// SetEnvDefault
func SetEnvDefault(key string, value string) {
	_, ok := os.LookupEnv(key)
	if ok {
		return
	}

	os.Setenv(key, value)
}
