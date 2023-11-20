package mongo

import (
	"nas-common/models"
	"nas-web/interal/db"

	"github.com/globalsign/mgo/bson"
	"github.com/kataras/iris/v12/context"
)

type chat struct{}

var Chat chat

func (chat) GetByUSid(ctx context.Context, usid string) (sessionMessagesDesc models.SessionMessagesDesc, err error, exist bool) {
	dbName := sessionMessagesDesc.Collection()
	query := bson.M{}
	query["_id"] = usid
	exist, err = db.MongoCli.FindOne(dbName, query, &sessionMessagesDesc)
	return
}

func (chat) AddSession(ctx context.Context, sessionMessagesDesc *models.SessionMessagesDesc) (err error) {
	dbName := sessionMessagesDesc.Collection()
	err = db.MongoCli.Insert(dbName, sessionMessagesDesc)
	return
}

func (chat) AppendMessages(ctx context.Context, usid string, update bson.M) (err error) {
	dbName := (&models.SessionMessagesDesc{}).Collection()
	query := bson.M{
		"_id": usid,
	}
	err = db.MongoCli.UpdateManual(dbName, query, update, false)
	return
}

func (chat) DeleteSession(ctx context.Context, usid string) (err error) {
	dbName := (&models.SessionMessagesDesc{}).Collection()
	query := bson.M{
		"_id": usid,
	}
	err = db.MongoCli.RemoveAll(dbName, query)
	return
}

func (chat) GetAllSessionsByUid(ctx context.Context, uid int) (sessionMessagesDescs []models.SessionMessagesDesc, err error) {
	dbName := (&models.SessionMessagesDesc{}).Collection()
	query := bson.M{
		"uid": uid,
	}
	err = db.MongoCli.FindAll(dbName, query, &sessionMessagesDescs)
	return
}
