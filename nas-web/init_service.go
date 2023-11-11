package nasweb

import (
	"nas-web/config"
	"nas-web/interal/cache"
	"nas-web/interal/db"
	"nas-web/interal/openai"
	"nas-web/middleware/jwt"
)

func InitService() bool {
	//init config file
	if err := config.ConfigInit("../confile/webapi_srv.yaml"); err != nil {
		panic(err)
	}
	//init mongo database
	db.MongoInit(config.IrisConfig.Mongodb)
	//init redis
	cache.RedisInit(config.IrisConfig.Redis)
	//init system token (avoid jwt redown)
	jwt.InitSysToken()
	//init openai
	openai.OpenaiInit(config.IrisConfig.Openai)
	return true
}
