//---------------------------------
//File Name    : ../support/deny_req.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2022-01-04 22:07:44
//Description  :
//----------------------------------
package sync

import (
	"nas-common/mlog"
	"nas-web/interal/cache"

	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

type SyncForce struct {
	c chan struct{}
	isSuccess bool
}

func (s *SyncForce) ForceRespToWait(channle string){
	s.c = make(chan struct{})
	//"1"-success | "0"-false
	err := cache.RedisCli.Subscribe(func(m redis.Message) error {
		s.isSuccess = func (d []byte)bool{
			if string(d) == "1"{
				return true
			}
			return false
		}(m.Data)
		close(s.c)
		return nil
	}, s.c, channle)
	if err != nil {
		mlog.Error("sync wait failed",zap.Error(err))
		s.isSuccess = false
	}
}

func (s *SyncForce) IsOk() bool{
	return s.isSuccess
}

func OpenWait(channel string)bool{
	var syncForce SyncForce
	syncForce.ForceRespToWait(channel)
	return syncForce.IsOk()
}
