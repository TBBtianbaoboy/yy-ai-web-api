package controller

import (
	formjson "nas-web/dao/form_json"
	"nas-web/interal/wrapper"
	"nas-web/service"
	"nas-web/support"
)

type UserController struct{}

// AddUser
// @Summary 添加用户
// @Description add user
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param auth body formjson.AddUserReq true "request data"
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /v1/user/ [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (u UserController) AddUser(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.AddUserHandler, true, &formjson.AddUserReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// DeleteUser
// @Summary 删除用户
// @Description delete user
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param auth body formjson.DeleteUserReq true "request data"
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /v1/user/ [delete]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (u UserController) DeleteUser(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.DeleteUserHandler, true, &formjson.DeleteUserReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// EditUser
// @Summary 编辑用户
// @Description edit user
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param auth body formjson.EditUserReq true "request data"
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /v1/user/ [put]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (u UserController) EditUser(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.EditUserHandler, true, &formjson.EditUserReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// ListUser
// @Summary 获取用户列表
// @Description list user
// @Tags 用户管理
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth query formjson.ListUserReq true "request data"
// @Success 200 {object} formjson.ListUserResp "response data"
// @Router /v1/user/ [get]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (u UserController) ListUser(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.ListUserHandler, true, &formjson.ListUserReq{PageSize: 10}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

// GetUserInfo
// @Summary 获取当前用户信息
// @Description get user info
// @Tags 用户管理
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} formjson.GetUserInfoResp "response data"
// @Router /v1/user/info [get]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (u UserController) GetUserInfo(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.GetUserInfoHandler, false, nil, nil)
}

// UpdateUserPasswd
// @Summary 修改用户密码
// @Description update user password
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param auth body formjson.UpdateUserPasswdReq true "request data"
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /v1/user/passwd/ [put]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (u UserController) UpdateUserPasswd(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.UpdateUserPasswdHandler, true, &formjson.UpdateUserPasswdReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// UpdateUserStatus
// @Summary 修改用户登录状态（是否允许登录）
// @Description update user login status
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param auth body formjson.UpdateUserStatusReq true "request data"
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /v1/user/status/ [put]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (u UserController) UpdateUserStatus(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.UpdateUserStatusHandler, true, &formjson.UpdateUserStatusReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// ResetPasswd
// @Summary 重置用户密码
// @Description reset user password
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param auth body formjson.ResetPasswdReq true "request data"
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /v1/user/reset_passwd [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (u UserController) ResetPasswd(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.ResetPasswdHandler, true, &formjson.ResetPasswdReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}
