//---------------------------------
//File Name    : jwt.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-17 16:25:46
//Description  :
//----------------------------------
package redis

import (
	"nas-web/interal/cache"
	"strconv"
)

//@Func 判断jwt是否在黑名单中
func IsJwtInBlackList(token string) bool {
	_,err := cache.RedisCli.Get(cache.JwtBlackList(token))
	if err != nil {
		return false
	}
	return true
}

//@Func 判断jwt是否在白名单中
func IsJwtInWhiteList(token string) bool {
	_,err := cache.RedisCli.Get(cache.JwtWhiteList(token))
	if err != nil {
		return false
	}
	return true
}

//@Func 刷新redis白名单中的jwt使用限时
func FlushJwtWhiteList(token string)(error){
	expire,err := cache.RedisCli.Get(cache.JwtWhiteList(token))
	if err != nil{
		return err
	}
	exp, _ := strconv.Atoi(expire)
	_, err = cache.RedisCli.SetEx(cache.JwtWhiteList(token), exp, int32(exp))
	return nil
}

//@Func 设置jwt token 白名单
func SetJwtWhiteList(token string, expire int32) (err error) {
	_, err = cache.RedisCli.SetEx(cache.JwtWhiteList(token), expire, expire)
	return
}

//@Func 设置jwt token 黑名单
func SetJwtBlacklist(token string, expire int32) (err error){
	_, err = cache.RedisCli.SetEx(cache.JwtBlackList(token),1, expire)
	return
}
