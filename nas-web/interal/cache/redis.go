//---------------------------------
//File Name    : interal/cache/redis.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-14 14:25:47
//Description  :
//----------------------------------
package cache

import (
	"nas-common/mredis"
	"nas-common/utils"
	"nas-web/config"
	"nas-web/support"
	"strings"
)

var RedisCli *mredis.RedisCli

func RedisInit(conf config.RedisConfig) {
	defaultCfg := mredis.NewDefaultOption()
	redisUrl := make([]string,0)
	redisUrl = append(redisUrl, conf.Host)
	redisUrl = append(redisUrl, utils.IntToString(conf.Port))
	defaultCfg.URL = strings.Join(redisUrl,":")
	defaultCfg.Password = conf.Passwd
	RedisCli = mredis.NewRedisCli(defaultCfg)
	if RedisCli == nil {
		panic(support.InitRedisError)
	}
}
