package mongo

import (
	"nas-common/models"
	"nas-web/interal/db"
	webutils "nas-web/web-utils"

	"github.com/globalsign/mgo/bson"
	"github.com/kataras/iris/v12/context"
)

type baseline struct{}

var Baseline baseline

func (baseline) Find(ctx context.Context, agentId string, count int, baselineDoc *models.CisScanOutline) (exist bool, err error) {
	dbName := baselineDoc.Collection()
	query := bson.M{
		"agent_id": agentId,
		"id":       count,
	}
	exist, err = db.MongoCli.FindOne(dbName, query, baselineDoc)
	return
}

func (baseline) FindCount(ctx context.Context, agentId string) (count int) {
	dbName := (&models.CisScanOutline{}).Collection()
	query := bson.M{
		"agent_id": agentId,
	}
	count, _ = db.MongoCli.FindCount(dbName, query)
	return
}

func (baseline) FindResult(ctx context.Context, query bson.M, page, pageSize int, scanResultDoc *[]models.CisScanResultItem) (count int, err error) {
	dbName := (&models.CisScanResultItem{}).Collection()
	limit := pageSize
	skip := webutils.GetPageStart(page, pageSize)
	err = db.MongoCli.FindByLimitAndSkip(dbName, query, scanResultDoc, limit, skip)
	count, _ = db.MongoCli.FindCount(dbName, query)
	return
}

func (baseline) FindOneResult(ctx context.Context, query bson.M, baselineDoc *models.CisScanResultItem) (exist bool, err error) {
	dbName := baselineDoc.Collection()
	exist, err = db.MongoCli.FindOne(dbName, query, baselineDoc)
	return
}

func (baseline) UpdateOutline(ctx context.Context, agentId string, id int, update bson.M) (err error) {
	dbName := (&models.CisScanOutline{}).Collection()
	query := bson.M{
		"agent_id": agentId,
		"id":       id,
	}
	err = db.MongoCli.Update(dbName, query, update, false)
	return
}

func (baseline) UpdateResult(ctx context.Context, query, update bson.M) (err error) {
	dbName := (&models.CisScanResultItem{}).Collection()
	err = db.MongoCli.Update(dbName, query, update, false)
	return
}

func (baseline) FindRepoByCisId(ctx context.Context, cisId string, doc *models.TbRepoCis) (err error) {
	dbName := (&models.TbRepoCis{}).Collection()
	query := bson.M{
		"id": cisId,
	}
	_, err = db.MongoCli.FindOne(dbName, query, doc)
	return
}
