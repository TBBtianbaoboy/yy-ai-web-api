//---------------------------------
//File Name    : system.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-17 18:14:10
//Description  :
//----------------------------------
package mongo

import (
	"nas-common/models"
	"nas-web/interal/db"

	"github.com/globalsign/mgo/bson"
	"github.com/kataras/iris/v12/context"
)

type systemConfig struct{}

var SystemConfig systemConfig

func (systemConfig) Get(ctx context.Context) (systemConfigDoc models.SystemConfig) {
	var ok bool
	dbName := systemConfigDoc.Collection()
	ok, _ = db.MongoCli.FindOne(dbName, bson.M{}, &systemConfigDoc)
	if !ok {
		systemConfigDoc = models.SystemConfig{
			PwdFlushCycle:     90,  //默认密码过期时间(天)
			NoOpExitTm:        300, //默认页面无操作自动退出时间(分钟)
			LoginFailedCount:  5,   //默认允许登录失败次数
			LoginFailedLockTm: 3,   //默认登录失败锁定时间(分钟)
		}
	}
	return
}
