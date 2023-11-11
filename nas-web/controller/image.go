package controller

import (
	formjson "nas-web/dao/form_json"
	"nas-web/interal/wrapper"
	"nas-web/service"
	"nas-web/support"
)

type AgentController struct{}

// ListAgent
// @Summary 获取agent列表
// @Description get agent list
// @Tags 资产管理
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth query formjson.ListAgentReq true "request data"
// @Success 200 {object} formjson.ListAgentResp "response data"
// @Router /v1/agent/ [get]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a AgentController) ListAgent(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.ListAgentHandler, true, &formjson.ListAgentReq{PageSize: 10}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

// AgentSystemInfo
// @Summary 获取agent的主机系统信息
// @Description get agent system info
// @Tags 资产管理
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth query formjson.AgentSystemInfoReq true "request data"
// @Success 200 {object} formjson.AgentSystemInfoResp "response data"
// @Router /v1/agent/info/system/ [get]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a AgentController) AgentSystemInfo(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.AgentSystemInfoHandler, true, &formjson.AgentSystemInfoReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

// AgentPortInfo
// @Summary 获取agent的主机对外开放端口信息
// @Description get agent port info
// @Tags 资产管理
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth query formjson.AgentPortInfoReq true "request data"
// @Success 200 {object} formjson.AgentPortInfoResp "response data"
// @Router /v1/agent/info/port/ [get]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a AgentController) AgentPortInfo(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.AgentPortInfoHandler, true, &formjson.AgentPortInfoReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

// DownloadAgent
// @Summary 下载agent
// @Description download agent
// @Tags 资产管理
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} formjson.DownloadAgentResp "response data"
// @Router /v1/agent/download/ [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a AgentController) DownloadAgent(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.DownloadAgentHandler, false, false, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

// DeleteAgent
// @Summary 删除agent
// @Description delete agent
// @Tags 资产管理
// @Accept json
// @Produce json
// @Param auth body formjson.DeleteAgentReq true "request data"
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /v1/agent/ [delete]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a AgentController) DeleteAgent(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.DeleteAgentHandler, true, &formjson.DeleteAgentReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// AddAgentSecGrpRule
// @Summary 增加agent安全组规则
// @Description add agent secure group rule
// @Tags 资产管理
// @Accept json
// @Produce json
// @Param auth body formjson.AddAgentSecGrpRuleReq true "request data"
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /v1/agent/secgrp [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a AgentController) AddAgentSecGrpRule(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.AddAgentSecGrpRuleHandler, true, &formjson.AddAgentSecGrpRuleReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// DeleteAgentSecGrpRule
// @Summary 删除agent安全组规则
// @Description delete agent secure group rule
// @Tags 资产管理
// @Accept json
// @Produce json
// @Param auth body formjson.DeleteAgentSecGrpRuleReq true "request data"
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /v1/agent/secgrp [delete]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a AgentController) DeleteAgentSecGrpRule(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.DeleteAgentSecGrpRuleHandler, true, &formjson.DeleteAgentSecGrpRuleReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// ListAgentSecGrp
// @Summary 获取agent 安全组列表
// @Description get agent secure group list
// @Tags 资产管理
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth query formjson.ListAgentSecGrpReq true "request data"
// @Success 200 {object} formjson.ListAgentSecGrpResp "response data"
// @Router /v1/agent/secgrp [get]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a AgentController) ListAgentSecGrp(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.ListAgentSecGrpHandler, true, &formjson.ListAgentSecGrpReq{PageSize: 10}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

// StartBaselineScan
// @Summary 开启基线扫描
// @Description start baseline scan
// @Tags 资产管理
// @Accept json
// @Produce json
// @Param auth body formjson.StartBaselineScanReq true "request data"
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /v1/agent/baseline [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a AgentController) StartBaselineScan(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.StartBaselineScanHandler, true, &formjson.StartBaselineScanReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// ListAgentBaseline
// @Summary 获取基线扫描结果列表
// @Description list agent baseline result
// @Tags 资产管理
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth query formjson.ListAgentBaselineReq true "request data"
// @Success 200 {object} formjson.ListAgentBaselineResp "response data"
// @Router /v1/agent/baseline [get]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a AgentController) ListAgentBaseline(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.ListAgentBaselineHandler, true, &formjson.ListAgentBaselineReq{PageSize: 10}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

// UpdateAgentBaseline
// @Summary 更新基线扫描项属性
// @Description update baseline scan item attr
// @Tags 资产管理
// @Accept json
// @Produce json
// @Param auth body formjson.UpdateAgentBaselineReq true "request data"
// @Success 200 {object} formjson.StatusResp "response data"
// @Router /v1/agent/baseline [put]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a AgentController) UpdateAgentBaseline(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.UpdateAgentBaselineHandler, true, &formjson.UpdateAgentBaselineReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// GetBaselineInfo
// @Summary 获取基线扫描详情
// @Description get baseline scan info
// @Tags 资产管理
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth query formjson.GetBaselineInfoReq true "request data"
// @Success 200 {object} formjson.GetBaselineInfoResp "response data"
// @Router /v1/agent/baseline/info [get]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a AgentController) GetBaselineInfo(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.GetBaselineInfoHandler, true, &formjson.GetBaselineInfoReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}
