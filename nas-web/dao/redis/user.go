//---------------------------------
//File Name    : user.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-17 18:24:41
//Description  :
//----------------------------------
package redis

import (
	"nas-web/interal/cache"
	"strconv"
)

//@Func 获取用户登陆失败的次数
func GetUserLoginLock(addr, username string) (count int) {
	countStr, _ := cache.RedisCli.Get(cache.UserLock(addr, username))
	count, _ = strconv.Atoi(countStr)
	return
}

//@Func 设置用户登陆失败的次数
func SetUserLoginLock(addr, username string, lockTm int32) {
	var count int
	countStr, _ := cache.RedisCli.Get(cache.UserLock(addr, username))
	count, _ = strconv.Atoi(countStr)
	count++
	_, _ = cache.RedisCli.SetEx(cache.UserLock(addr, username), count, lockTm*60)
}

//@Func 移除用户登陆失败的缓存记录
func RemoveUserLoginLock(addr, username string) {
	_ = cache.RedisCli.Del(cache.UserLock(addr, username))
}
