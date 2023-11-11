//---------------------------------
//File Name    : interal/cache/key.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-17 16:35:33
//Description  :
//----------------------------------
package cache

//@Func 获取reids jwt黑名单中的key
func JwtBlackList(token string) string {
	return "jwt:blacklist:" + token
}

//@Func 获取redis jwt白名单中的key
func JwtWhiteList(token string) string {
	return "jwt:whitelist:" + token
}

//@Func 获取redis 中的用户锁定表中的key
func UserLock(addr, username string) string {
	return "user:lock:" + addr + ":" + username
}

//@Func 获取redis中agent的key
func AgentKey(mac string) string {
	return "agent:" + mac
}

//@Func 获取baseline中同步响应key
func BaselineSync(channel string) string {
	return "baseline:" + channel
}
