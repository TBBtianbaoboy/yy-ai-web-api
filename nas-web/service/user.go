//---------------------------------
//File Name    : service/user.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-18 16:26:53
//Description  :
//----------------------------------
package service

import (
	"fmt"
	"nas-common/mlog"
	"nas-common/models"
	"nas-common/utils"
	formjson "nas-web/dao/form_json"
	"nas-web/dao/mongo"
	"nas-web/interal/password"
	"nas-web/interal/wrapper"
	"nas-web/support"
	webutils "nas-web/web-utils"
	"regexp"
	"time"

	"github.com/globalsign/mgo/bson"
	"go.uber.org/zap"
)

// AddUserHandler 新增用户
func AddUserHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.AddUserReq)
	resp := formjson.StatusResp{Status: "OK"}
	//判断当前用户是否有权限增添用户
	if ctx.UserToken.UserType != models.UserType_Admin {
		support.SendApiErrorResponse(ctx, support.NoPrivilegeAddUser, 0)
		return nil
	}
	//判断密码是否相同
	if !webutils.String.Compare(req.Password, req.Confirm) {
		support.SendApiErrorResponse(ctx, support.PasswordNotConfirm, 0)
		return nil
	}
	//判断用户是否存在
	existQuery := bson.M{"username": req.Username}
	if mongo.User.IsExist(ctx, existQuery) {
		support.SendApiErrorResponse(ctx, support.UserIsExist, 0)
		return nil
	}
	//判断密码强度
	passwordStrengthLevel := webutils.String.GetPasswordStrength(req.Password)
	if passwordStrengthLevel == 0 {
		support.SendApiErrorResponse(ctx, support.PasswordStrengthFailed, 0)
		return nil
	}
	newUid := mongo.User.GetMaxUid(ctx)
	// 创建账户
	userDoc := models.User{
		ID:               bson.NewObjectId(),
		UID:              newUid,
		Enable:           true,
		UserType:         req.UserType,
		UserName:         req.Username,
		Password:         password.MakePassword(req.Password),
		PasswordStrength: passwordStrengthLevel,
		Mail:             req.Mail,
		Mobile:           req.Mobile,
		LastPwdChangeTm:  time.Now(),
		LastLoginTm:      time.Now(),
		InsertTm:         time.Now(),
		UpdateTm:         time.Now(),
	}

	if err = mongo.User.Create(ctx, userDoc); err != nil {
		mlog.Error("create user failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.CreateUserFailed, 0)
		return nil
	}

	support.SendApiResponse(ctx, resp, "")
	return nil
}

// DeleteUserHandler 删除用户(支持批量删除)
func DeleteUserHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.DeleteUserReq)
	resp := formjson.StatusResp{Status: "OK"}
	//只有管理员才可以删除用户
	if ctx.UserToken.UserType != models.UserType_Admin {
		mlog.Error("permission deny", zap.String("operate Uid", utils.IntToString(ctx.UserToken.UserId)))
		support.SendApiErrorResponse(ctx, support.PermissionDeny, 0)
		return nil
	}
	for _, id := range req.Uids {
		//管理员不可删除自己
		if id == 1 {
			continue
		}
		if err = mongo.User.Delete(ctx, id); err != nil {
			mlog.Error("delete user failed", zap.String("uid", utils.IntToString(id)))
		}
	}
	support.SendApiResponse(ctx, resp, "")
	return
}

// EditUserHandler 编辑用户 ok
func EditUserHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.EditUserReq)
	resp := formjson.StatusResp{Status: "OK"}
	query := bson.M{
		"uid": req.Uid,
	}
	//管理员任意编辑，(超级)用户只能编辑自己
	if ctx.UserToken.UserType != models.UserType_Admin && ctx.UserToken.UserId != req.Uid {
		mlog.Error("permission deny", zap.String("operate Uid", utils.IntToString(ctx.UserToken.UserId)))
		support.SendApiErrorResponse(ctx, support.PermissionDeny, 0)
		return nil
	}
	updateQuery := bson.M{
		"mail":   req.Mail,
		"mobile": req.Mobile,
	}
	if err = mongo.User.Edit(ctx, query, updateQuery); err != nil {
		mlog.Error("edit user failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.EditUserFailed, 0)
		return nil
	}
	support.SendApiResponse(ctx, resp, "")
	return
}

// ListUserHandler 获取用户列表
func ListUserHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.ListUserReq)
	resp := formjson.ListUserResp{}
	query := bson.M{}
	if req.Search != "" {
		query["$or"] = []bson.M{
			{"username": bson.M{"$regex": bson.RegEx{Pattern: regexp.QuoteMeta(req.Search), Options: "i"}}},
			{"mail": bson.M{"$regex": bson.RegEx{Pattern: regexp.QuoteMeta(req.Search), Options: "i"}}},
			{"mobile": bson.M{"$regex": bson.RegEx{Pattern: regexp.QuoteMeta(req.Search), Options: "i"}}},
		}
	}
	if req.Enable == 1 {
		query["enable"] = true
	} else if req.Enable == -1 {
		query["enable"] = false
	}
	if req.UserType != 0 {
		query["user_type"] = req.UserType
	}
	var userDocs []models.User
	if resp.Count, userDocs, err = mongo.User.List(ctx, query, req.Page, req.PageSize); err != nil {
		mlog.Error("edit user failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ListUsersFailed, 0)
		return nil
	}

	for _, userDoc := range userDocs {
		resp.Results = append(resp.Results, formjson.ListUserItem{
			Uid:        userDoc.UID,
			Username:   userDoc.UserName,
			UserType:   userDoc.UserType,
			Enable:     userDoc.Enable,
			Mail:       userDoc.Mail,
			Mobile:     userDoc.Mobile,
			CreateTime: userDoc.InsertTm.Unix(),
		})
	}
	support.SendApiResponse(ctx, resp, "")
	return
}

// UpdateUserStatus 更新用户登录状态 ok
func UpdateUserStatusHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.UpdateUserStatusReq)
	resp := formjson.StatusResp{Status: "OK"}
	//只有管理员可以修改登录状态
	if ctx.UserToken.UserType != models.UserType_Admin {
		mlog.Error("permission deny", zap.String("operate Uid", utils.IntToString(ctx.UserToken.UserId)))
		support.SendApiErrorResponse(ctx, support.PermissionDeny, 0)
		return nil
	}
	query := bson.M{
		"uid": req.Uid,
	}
	if isExist := mongo.User.IsExist(ctx, query); !isExist {
		mlog.Error("user not exist", zap.String("Uid", utils.IntToString(req.Uid)))
		support.SendApiErrorResponse(ctx, support.UserNotExist, 0)
		return nil
	}
	updateQuery := bson.M{
		"enable": !req.Enable,
	}
	if err = mongo.User.Edit(ctx, query, updateQuery); err != nil {
		mlog.Error("update user status failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.UpdateUserStatusFailed, 0)
		return nil
	}
	support.SendApiResponse(ctx, resp, "")
	return
}

//GetUserInfoHandler 获取当前用户个人信息 ok
func GetUserInfoHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	var userDoc models.User
	if err = mongo.User.FindByUid(ctx, ctx.UserToken.UserId, &userDoc); err != nil {
		mlog.Error("get current user info failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.GetUserInfoFailed, 0)
		return nil
	}
	resp := formjson.GetUserInfoResp{
		Uid:        userDoc.UID,
		UserType:   userDoc.UserType,
		UserName:   userDoc.UserName,
		Email:      userDoc.Mail,
		Phone:      userDoc.Mobile,
		PS:         userDoc.PasswordStrength,
		CreateTime: userDoc.InsertTm.Unix(),
	}
	fmt.Println(userDoc.Mail)
	support.SendApiResponse(ctx, resp, "")
	return
}

// UpdateUserPasswdHandler 更新用户密码 ok
func UpdateUserPasswdHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.UpdateUserPasswdReq)
	resp := formjson.StatusResp{Status: "OK"}
	//所有用户都只能修改自己的密码
	if ctx.UserToken.UserId != req.Uid {
		support.SendApiErrorResponse(ctx, support.PermissionDeny, 0)
		return nil
	}

	if req.New != req.New2 {
		support.SendApiErrorResponse(ctx, support.PasswordDiffFailed, 0)
		return nil
	}

	var userDoc models.User
	if err = mongo.User.FindByUid(ctx, req.Uid, &userDoc); err != nil {
		support.SendApiErrorResponse(ctx, support.GetUserInfoFailed, 0)
		return nil
	}

	if !password.CheckPassword(req.Old, userDoc.Password) {
		support.SendApiErrorResponse(ctx, support.PasswordFailed, 0)
		return nil
	}

	query := bson.M{
		"uid": req.Uid,
	}
	update := bson.M{
		"password": password.MakePassword(req.New),
	}
	if err = mongo.User.Edit(ctx, query, update); err != nil {
		support.SendApiErrorResponse(ctx, support.UpdateUserPasswdFailed, 0)
		return nil
	}

	support.SendApiResponse(ctx, resp, "")
	return
}

// ResetPasswdHandler 重置用户密码
func ResetPasswdHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.ResetPasswdReq)
	resp := formjson.StatusResp{Status: "OK"}
	//权限控制 TODO 可以将所有权限控制优化到中间件中
	if ctx.UserToken.UserType != models.UserType_Admin || (ctx.UserToken.UserType == models.UserType_Admin && req.Uid == ctx.UserToken.UserId) {
		support.SendApiErrorResponse(ctx, support.PermissionDeny, 0)
		return nil
	}
	if exist := mongo.User.IsExist(ctx, bson.M{"uid": req.Uid}); !exist {
		support.SendApiErrorResponse(ctx, support.UserNotExist, 0)
		return nil
	}

	query := bson.M{
		"uid": req.Uid,
	}
	update := bson.M{
		"password": password.MakePassword(req.Password),
	}
	if err = mongo.User.Edit(ctx, query, update); err != nil {
		support.SendApiErrorResponse(ctx, support.UpdateUserPasswdFailed, 0)
		return nil
	}

	support.SendApiResponse(ctx, resp, "")
	return
}
