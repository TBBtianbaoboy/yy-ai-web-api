//---------------------------------
//File Name    : auth.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-14 17:23:17
//Description  :
//----------------------------------
package controller

import (
	formjson "nas-web/dao/form_json"
	"nas-web/interal/wrapper"
	"nas-web/service"
	"nas-web/support"
)

type AuthController struct {
}

// VerifyCode
// @Summary 基础接口 - 获取验证码
// @Description get verifycode
// @Tags common
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} formjson.VerifyCodeResp "response data"
// @Router /auth/verifycode/ [get]
// @Security ApiKeyAuth
func (a AuthController) VerifyCode(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.VerifyCodeHandler, false, nil, nil)
}

// Login
// @Summary 基础接口 - 用户登录
// @Description user login
// @Tags common
// @Accept json
// @Produce json
// @Param auth body formjson.LoginReq true "request data"
// @Success 200 {object} formjson.LoginResp "response data"
// @Router /auth/login/ [post]
// @Security ApiKeyAuth
func (a AuthController) Login(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.LoginHandler, true,&formjson.LoginReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// Logout
// @Summary 基础接口 - 用户登出
// @Description user logout
// @Tags common
// @Accept json
// @Produce json
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /auth/logout/ [post]
// @Security ApiKeyAuth
// @Param authorization header string true "authorization"
func (a AuthController) Logout(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.LogoutHandler, false, nil, nil)
}
