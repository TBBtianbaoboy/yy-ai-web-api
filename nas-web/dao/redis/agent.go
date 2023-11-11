//---------------------------------
//File Name    : dao/redis/agent.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-22 02:09:14
//Description  :
//----------------------------------
package redis

import "nas-web/interal/cache"

//@Func 判断agent是否存活
func IsAgentLive(mac string) (exist bool) {
	exist,_ = cache.RedisCli.Exists(cache.AgentKey(mac))
	return
}
