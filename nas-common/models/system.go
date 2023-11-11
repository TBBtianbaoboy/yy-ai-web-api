//---------------------------------
//File Name    : system.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-17 18:15:57
//Description  : 
//----------------------------------
package models

type SystemConfig struct {
	PwdFlushCycle     int  `bson:"pwd_flush_cycle"`      //密码更新间隔
	NoOpExitTm        int  `bson:"no_oper_exit_tm"`      //无操作退出时间
	LoginFailedCount  int  `bson:"login_failed_count"`   //允许登录失败次数
	LoginFailedLockTm int  `bson:"login_failed_lock_tm"` //登录失败锁定时间
}

func (s *SystemConfig)Collection()string {
	return "tb_system_config"
}
