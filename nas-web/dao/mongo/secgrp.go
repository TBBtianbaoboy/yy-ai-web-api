//---------------------------------
//File Name    : secgrp.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-31 10:10:58
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

type secGrp struct{}

var SecGrp secGrp

func (secGrp) Find(ctx context.Context,query bson.M ,secGrpDoc *models.SecGrp) (exist bool){
	dbName := (&models.SecGrp{}).Collection()
	exist,_ = db.MongoCli.FindOne(dbName,query,secGrpDoc)
	return
}

func (secGrp) Insert(ctx context.Context,secGrpDoc *models.SecGrp)(err error){
	dbName := secGrpDoc.Collection()
	err = db.MongoCli.Insert(dbName,secGrpDoc)
	return
}

func (secGrp) Delete(ctx context.Context,query bson.M)(err error){
	dbName := (&models.SecGrp{}).Collection()
	err = db.MongoCli.RemoveAll(dbName,query)
	return
}

func (secGrp) List(ctx context.Context,query bson.M,secGrpDocs *[]models.SecGrp,page,pageSize int,sortList ...string)(count int,err error){
	dbName := (&models.SecGrp{}).Collection()
	limit := pageSize
	skip := webutils.GetPageStart(page, pageSize)
	if err = db.MongoCli.FindSortByLimitAndSkip(dbName,query,secGrpDocs,limit,skip,sortList...);err != nil {
		return
	}
	count,err = db.MongoCli.FindCount(dbName,query)
	return
}
