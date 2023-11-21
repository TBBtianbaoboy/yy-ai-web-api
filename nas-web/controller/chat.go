package controller

import (
	formjson "nas-web/dao/form_json"
	"nas-web/interal/wrapper"
	"nas-web/service"
	"nas-web/support"
)

type ChatController struct{}

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
// @Router /v1/chat/get_all_sessions [get]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a ChatController) GetAllSessions(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.GetAllSessionsHandler, false, nil, nil)
}

// @Summary 获取指定会话的所有消息
// @Description get session messages by session id
// @Tags Chat
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth query formjson.GetSessionMessagesReq true "request data"
// @Success 200 {object} formjson.GetSessionMessagesResp "response data"
// @Router /v1/chat/get_session_messages [get]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a ChatController) GetSessionMessages(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.GetSessionMessagesHandler, true, &formjson.GetSessionMessagesReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

// @Summary 新建会话
// @Description create new session
// @Tags Chat
// @Accept json
// @Produce json
// @Param auth body formjson.CreateSessionReq true "request data"
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /v1/chat/session [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a ChatController) CreateSession(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.CreateSessionHandler, true, &formjson.CreateSessionReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// @Summary 更新会话
// @Description update existed session
// @Tags Chat
// @Accept json
// @Produce json
// @Param auth body formjson.UpdateSessionReq true "request data"
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /v1/chat/session [put]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a ChatController) UpdateSession(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.UpdateSessionHandler, true, &formjson.UpdateSessionReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}
