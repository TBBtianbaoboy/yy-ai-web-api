//---------------------------------
//File Name    : user.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-17 17:50:21
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

type user struct{}

var User user

func (user) FindByName(ctx context.Context, name string) (userDoc models.User, err error) {
	dbName := userDoc.Collection()
	query := bson.M{"username": name}
	_, err = db.MongoCli.FindOne(dbName, query, &userDoc)
	return
}

func (user) FindByUid(ctx context.Context, uid int, userDoc *models.User) (err error) {
	dbName := userDoc.Collection()
	query := bson.M{"uid": uid}
	_, err = db.MongoCli.FindOne(dbName, query, userDoc)
	return
}

func (user) IsExist(ctx context.Context, query bson.M) bool {
	var userDoc []models.User
	dbName := (&models.User{}).Collection()
	if err := db.MongoCli.FindAll(dbName, query, &userDoc); err != nil || len(userDoc) == 0 {
		return false
	}
	return true
}

func (user) GetMaxUid(ctx context.Context) (uid int) {
	var userDoc []models.User
	dbName := (&models.User{}).Collection()
	_ = db.MongoCli.FindSortByLimitAndSkip(dbName, bson.M{}, &userDoc, 1, 0, "-uid")
	uid = userDoc[0].UID + 1
	return
}

func (user) Create(ctx context.Context, user models.User) (err error) {
	dbName := user.Collection()
	err = db.MongoCli.Insert(dbName, user)
	return
}

func (user) List(ctx context.Context, query bson.M, page, pageSize int) (count int, userDocs []models.User, err error) {
	dbName := (&models.User{}).Collection()
	skip := webutils.GetPageStart(page, pageSize)
	err = db.MongoCli.FindByLimitAndSkip(dbName, query, &userDocs, pageSize, skip)
	count, _ = db.MongoCli.FindCount(dbName, query)
	return
}

func (user) Delete(ctx context.Context, uid int) (err error) {
	dbName := (&models.User{}).Collection()
	query := bson.M{
		"uid": uid,
	}
	err = db.MongoCli.RemoveAll(dbName, query)
	return
}

func (user) Edit(ctx context.Context, query bson.M, update bson.M) (err error) {
	dbName := (&models.User{}).Collection()
	err = db.MongoCli.Update(dbName, query, update, false)
	return
}

func (user) AddLoginHistory(ctx context.Context, userLoginHistoryDoc *models.UserLoginHistory) (err error) {
	dbName := userLoginHistoryDoc.Collection()
	err = db.MongoCli.Insert(dbName, userLoginHistoryDoc)
	return
}
