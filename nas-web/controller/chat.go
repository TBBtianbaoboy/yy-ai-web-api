package controller

import (
	formjson "nas-web/dao/form_json"
	"nas-web/interal/wrapper"
	"nas-web/service"
	"nas-web/support"
)

type ChatController struct{}

// @Summary 发送无上下文无流式聊天
// @Description send no context no stream chat
// @Tags Chat
// @Accept json
// @Produce json
// @Param auth body formjson.SendNoContextNoStreamChatReq true "request data"
// @Success 200 {object} formjson.SendNoContextNoStreamChatResp "response data"
// @Router /v1/chat/no_context_no_stream [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a ChatController) SendNoContextNoStreamChat(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.SendNoContextNoStreamChatHandler, true, &formjson.SendNoContextNoStreamChatReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// @Summary 发送无上下文流式聊天
// @Description send no context stream chat
// @Tags Chat
// @Accept json
// @Produce json
// @Param auth body formjson.SendNoContextStreamChatReq true "request data"
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /v1/chat/no_context_stream [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a ChatController) SendNoContextStreamChat(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.SendNoContextStreamChatHandler, true, &formjson.SendNoContextStreamChatReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// @Summary 发送上下文流式聊天
// @Description send support context stream chat
// @Tags Chat
// @Accept json
// @Produce json
// @Param auth body formjson.SendContextStreamChatReq true "request data"
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /v1/chat/context_stream [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a ChatController) SendContextStreamChat(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.SendContextStreamChatHandler, true, &formjson.SendContextStreamChatReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// @Summary 删除上下文流式聊天
// @Description delete context stream chat
// @Tags Chat
// @Accept json
// @Produce json
// @Param auth body formjson.DeleteContextStreamChatReq true "request data"
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /v1/chat/delete_context_stream [delete]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a ChatController) DeleteContextStreamChat(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.DeleteContextStreamChatHandler, true, &formjson.DeleteContextStreamChatReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// @Summary 获取指定用户的所有会话列表
// @Description get all sessions list by user id
// @Tags Chat
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} formjson.GetAllSessionsResp "response data"
// @Router /v1/chat/get_sessions [get]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a ChatController) GetAllSessions(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.GetAllSessionsHandler, false, nil, nil)
}
