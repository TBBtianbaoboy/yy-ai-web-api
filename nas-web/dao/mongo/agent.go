//---------------------------------
//File Name    : dao/mongo/agent.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-22 01:43:03
//Description  :
//----------------------------------
package mongo

import (
	"nas-common/models"
	"nas-web/interal/db"
	webutils "nas-web/web-utils"

	"github.com/globalsign/mgo/bson"
	"github.com/kataras/iris/v12/context"
)

type agent struct{}

var Agent agent

func (agent)List(ctx context.Context,query bson.M,page,pageSize int)(count int,agentInfoDocs []models.AgentHandlerInfo,err error){
	dbName := (&models.AgentHandlerInfo{}).Collection()
	limit := pageSize
	skip := webutils.GetPageStart(page, pageSize)
	err = db.MongoCli.FindByLimitAndSkip(dbName, query, &agentInfoDocs, limit, skip)
	count, _ = db.MongoCli.FindCount(dbName, query)
	return
}

func (agent)Delete(ctx context.Context,query bson.M) (err error){
	dbName := (&models.AgentHandlerInfo{}).Collection()
	err = db.MongoCli.RemoveAll(dbName,query)
	return
}

//--------------------------------------------------------------

type agentDetails struct {}

var AgentDetails agentDetails

func (agentDetails)Get(ctx context.Context,query bson.M)(agentDetailsDoc models.AgentHandlerInfoDetails,err error){
	dbName := agentDetailsDoc.Collection()
	_,err = db.MongoCli.FindOne(dbName,query,&agentDetailsDoc)
	return
}
