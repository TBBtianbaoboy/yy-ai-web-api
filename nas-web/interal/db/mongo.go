package db

import (
	"fmt"
	"nas-common/mdb"
	"nas-common/mlog"
	"nas-web/config"

	"go.uber.org/zap"
)

var MongoCli mdb.DBAdaptor

func MongoInit(conf config.MongodbConfig) {
	var mongoURL = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authMechanism=SCRAM-SHA-1&authSource=admin", conf.User, conf.Passwd, conf.Host, conf.Port, conf.DbName)
	MongoCli = mdb.NewMongoSession()
	err := MongoCli.Connect(mongoURL)
	if err != nil {
		mlog.Error("Connect Mongodb failed", zap.Error(err), zap.String("url", mongoURL))
		panic(err)
	}
}
