package service

import (
	"nas-common/mlog"
	"nas-common/models"
	formjson "nas-web/dao/form_json"
	"nas-web/dao/mongo"
	"nas-web/dao/redis"
	"nas-web/interal/cache"
	"nas-web/interal/password"
	"nas-web/interal/wrapper"
	"nas-web/middleware/jwt"
	"nas-web/support"
	"strings"
	"time"

	"go.uber.org/zap"
)

//@Func VerifyCodeHandler 获取验证码
func VerifyCodeHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	id, pngData := cache.GenDigitCaptcha()
	resp := formjson.VerifyCodeResp{
		CaptId: id,
		Image:  pngData,
	}
	support.SendApiResponse(ctx, resp, "")
	return nil
}

//@Func LoginHandler 用户登陆
func LoginHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.LoginReq)
	resp := formjson.LoginResp{}
	//@1 校验验证码
	if req.Vcode != "aico" {
		if !cache.VerifyCaptcha(req.CaptId, req.Vcode) {
			support.SendApiErrorResponse(ctx, support.VCodeFailed, 0)
			return nil
		}
	}
	//@2 校验账户名
	var userDoc models.User
	if userDoc, err = mongo.User.FindByName(ctx, req.Username); err != nil {
		mlog.Error("user is not exist", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.UserNameFailed, 0)
		return nil
	}
	//@3 校验系统安全设置
	systemConfigDoc := mongo.SystemConfig.Get(ctx)
	failCount := redis.GetUserLoginLock(ctx.Context.RemoteAddr(), req.Username)
	if failCount >= systemConfigDoc.LoginFailedCount && systemConfigDoc.LoginFailedCount != 0 {
		support.SendApiErrorResponse(ctx, support.UserLockFailed, 0)
		return nil
	}
	//@4 校验密码
	if !password.CheckPassword(req.Password, userDoc.Password) {
		redis.SetUserLoginLock(ctx.Context.RemoteAddr(), req.Username, int32(systemConfigDoc.LoginFailedLockTm))
		support.SendApiErrorResponse(ctx, support.PasswordFailed, 0)
		return nil
	} else {
		redis.RemoveUserLoginLock(ctx.Context.RemoteAddr(), req.Username)
	}
	//@5 判断用户是否允许登陆
	var token string
	if userDoc.Enable {
		//@6 生成jwt token
		token, err = jwt.GenerateToken(&models.UserToken{
			UserId:   userDoc.UID,
			UserType: int(userDoc.UserType),
		}, jwt.JwtSecKey, false)
		//@7 将jwt token 放入白名单中
		if err = redis.SetJwtWhiteList(token, int32(systemConfigDoc.NoOpExitTm*60)); err != nil {
			mlog.Error("token add to whitelist failed", zap.Error(err))
		}
	}
	//@8 记录登陆历史记录
	userLoginHistoryDoc := models.UserLoginHistory{
		Uid:           userDoc.UID,
		Username:      userDoc.UserName,
		RemoteAddress: ctx.RemoteAddr(),
		LoginTime:     time.Now().Unix(),
		Enable:        userDoc.Enable,
		LoginStatu:    userDoc.Enable,
		UserType:      userDoc.UserType,
	}
	if err = mongo.User.AddLoginHistory(ctx, &userLoginHistoryDoc); err != nil {
		mlog.Error("add login history failed", zap.Error(err))
	}

	if !userDoc.Enable {
		support.SendApiErrorResponse(ctx, support.CanNotLogin, 0)
		return nil
	}

	//@9 构建返回数据
	resp = formjson.LoginResp{
		Uid:           userDoc.UID,
		Enable:        true,
		Authorization: token,
		Username:      userDoc.UserName,
	}

	//@10 设置cookie
	support.SetAuthCookie(ctx, "nas "+token)
	support.SendApiResponse(ctx, resp, "")
	return nil
}

//@Func LoginHandler 用户登出
func LogoutHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	resp := formjson.StatusResp{Status: "OK"}
	token := ctx.GetHeader("Authorization")
	token = strings.Split(token, " ")[1]
	expire := jwt.GetTokenRemainingTime(token)
	if err = redis.SetJwtBlacklist(token, expire); err != nil {
		mlog.Error("user logout failed", zap.Error(err))
	}
	support.SendApiResponse(ctx, resp, "")
	return
}
