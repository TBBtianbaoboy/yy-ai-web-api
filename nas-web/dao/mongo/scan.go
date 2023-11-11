package mongo

import (
	"nas-common/models"
	"nas-web/interal/db"
	webutils "nas-web/web-utils"

	"github.com/globalsign/mgo/bson"
	"github.com/kataras/iris/v12/context"
)

type scan struct{}

var Scan scan

func (scan) Add(ctx context.Context, scanStatusDescDoc *models.ScanStatusDesc) (err error) {
	dbName := scanStatusDescDoc.Collection()
	err = db.MongoCli.Insert(dbName, scanStatusDescDoc)
	return
}

func (scan) List(ctx context.Context, query bson.M, page, pageSize int, sort ...string) (count int, scanStatusDescDoc []models.ScanStatusDesc, err error) {
	dbName := (&models.ScanStatusDesc{}).Collection()
	limit := pageSize
	skip := webutils.GetPageStart(page, pageSize)
	err = db.MongoCli.FindSortByLimitAndSkip(dbName, query, &scanStatusDescDoc, limit, skip, sort...)
	count, _ = db.MongoCli.FindCount(dbName, query)
	return
}

func (scan) GetByScanID(ctx context.Context, scanId string) (err error, scanDoc models.ScanStatusDesc) {
	dbName := (&models.ScanStatusDesc{}).Collection()
	query := bson.M{}
	query["_id"] = scanId
	_, err = db.MongoCli.FindOne(dbName, query, &scanDoc)
	return
}

func (scan) Delete(ctx context.Context, scanId string) (err error) {
	dbName := (&models.ScanStatusDesc{}).Collection()
	query := bson.M{
		"_id": scanId,
	}
	err = db.MongoCli.RemoveAll(dbName, query)
	return
}

func (scan) GetFirstScanByID(ctx context.Context, query bson.M) (err error, doc models.DeepScanIpResult) {
	dbName := doc.Collection()
	_, err = db.MongoCli.FindOne(dbName, query, &doc)
	return
}

func (scan) GetFirstScanCountByID(ctx context.Context, query bson.M) (err error, doc models.FirstScanIpResult) {
	dbName := doc.Collection()
	_, err = db.MongoCli.FindOne(dbName, query, &doc)
	return
}
