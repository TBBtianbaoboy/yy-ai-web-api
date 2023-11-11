package service

import (
	"io/ioutil"
	"nas-common/mlog"
	"nas-common/models"
	"nas-common/rpcapi/forward"
	"nas-common/utils"
	formjson "nas-web/dao/form_json"
	"nas-web/dao/mongo"
	"nas-web/dao/redis"
	"nas-web/interal/cache"
	"nas-web/interal/sync"
	"nas-web/interal/wrapper"
	"nas-web/pkg/rpc"
	"nas-web/support"
	webutils "nas-web/web-utils"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

// ListAgentHandler
func ListAgentHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.ListAgentReq)
	resp := formjson.ListAgentResp{}
	query := bson.M{"is_deleted": false}
	if req.Search != "" {
		query["$or"] = []bson.M{
			{"agent_ip": bson.M{"$regex": bson.RegEx{Pattern: regexp.QuoteMeta(req.Search), Options: "i"}}},
			{"hostname": bson.M{"$regex": bson.RegEx{Pattern: regexp.QuoteMeta(req.Search), Options: "i"}}},
		}
	}
	var agentInfoDocs []models.AgentHandlerInfo
	if resp.Count, agentInfoDocs, err = mongo.Agent.List(ctx, query, req.Page, req.PageSize); err != nil {
		mlog.Error("list agent info failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ListAgentInfoFailed, 0)
		return nil
	}
	for _, agentInfo := range agentInfoDocs {
		resp.Results = append(resp.Results, formjson.ListAgentItem{
			HashId:     agentInfo.HashId,
			AgentIp:    agentInfo.AgentIp,
			Hostname:   agentInfo.Hostname,
			JoinTime:   agentInfo.Jointime,
			Pid:        agentInfo.Pid,
			UpdateTime: agentInfo.UpdateTime,
			Status:     redis.IsAgentLive(agentInfo.MacAddress),
		})
	}
	support.SendApiResponse(ctx, resp, "")
	return
}

// AgentSystemInfoHandler
func AgentSystemInfoHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.AgentSystemInfoReq)
	query := bson.M{
		"hash_id": req.HashId,
	}

	var agentInfoDetails models.AgentHandlerInfoDetails
	if agentInfoDetails, err = mongo.AgentDetails.Get(ctx, query); err != nil {
		mlog.Error("get agent info details failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.GetAgentInfoDetailsFailed, 0)
		return nil
	}

	resp := formjson.AgentSystemInfoResp{
		CpuUsed:     webutils.Decimal(agentInfoDetails.Resources.CpuUsed),
		MemoryUsed:  webutils.Decimal(agentInfoDetails.Resources.MemoryUsed),
		CpuCore:     agentInfoDetails.Resources.CpuCore,
		MemoryTotal: webutils.FormatSize(agentInfoDetails.Resources.MemoryTotal),
		DiskInfo: formjson.AgentSystemInfoDisk{
			Device:      agentInfoDetails.DiskInfo.Device,
			Fstype:      agentInfoDetails.DiskInfo.Fstype,
			MountPoint:  agentInfoDetails.DiskInfo.MountPoint,
			Options:     agentInfoDetails.DiskInfo.Options,
			Total:       webutils.FormatSize(agentInfoDetails.DiskInfo.Total),
			UsedPercent: webutils.Decimal(agentInfoDetails.DiskInfo.UsedPercent),
		},
	}

	support.SendApiResponse(ctx, resp, "")
	return
}

// DownloadAgentHandler
func DownloadAgentHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	resp := formjson.DownloadAgentResp{Status: "OK"}
	path := "/home/aico/tmp/a.zip"
	file, err := os.Open(path)
	if err != nil {
		support.SendApiErrorResponse(ctx, "文件不存在", iris.StatusNotFound)
		return
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		support.SendApiErrorResponse(ctx, "读取文件失败", iris.StatusInternalServerError)
		return
	}
	ctx.Context.Header("Content-Disposition", "attachment; filename=agent-1.0.0.zip")
	ctx.Context.Header("Content-Type", "application/octet-stream")
	// ctx.Context.Header("Content-Length", utils.IntToString(len(content)))
	// ctx.Context.Header("Accept-Ranges", "bytes")
	ctx.Context.Write(content)

	support.SendApiResponse(ctx, resp, "")
	return
}

// AgentPortInfoHandler
func AgentPortInfoHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.AgentPortInfoReq)
	resp := formjson.AgentPortInfoResp{}
	query := bson.M{
		"hash_id": req.HashId,
	}

	//子数组模糊查询
	if req.Search != "" {
		// query["$or"] = []bson.M{
		//     {"port": bson.M{"$regex": bson.RegEx{Pattern: regexp.QuoteMeta(req.Search), Options: "i"}}},
		//     {"port_service": bson.M{"$regex": bson.RegEx{Pattern: regexp.QuoteMeta(req.Search), Options: "i"}}},
		//     {"port_infos": bson.M{"$elemMatch": bson.M{"pid": bson.M{"$regex": bson.RegEx{Pattern: regexp.QuoteMeta(req.Search), Options: "i"}},
		//     }}},
		// }
		query["$port_infos"] = bson.M{
			"$elemMatch": bson.M{
				"$or": []bson.M{
					{"port": bson.M{"$regex": bson.RegEx{Pattern: regexp.QuoteMeta(req.Search), Options: "i"}}},
					{"port_service": bson.M{"$regex": bson.RegEx{Pattern: regexp.QuoteMeta(req.Search), Options: "i"}}},
					{"port_infos": bson.M{"$elemMatch": bson.M{"pid": bson.M{"$regex": bson.RegEx{Pattern: regexp.QuoteMeta(req.Search), Options: "i"}}}}},
				}}}
	}

	var agentInfoDetails models.AgentHandlerInfoDetails
	if agentInfoDetails, err = mongo.AgentDetails.Get(ctx, query); err != nil {
		mlog.Error("get agent info details failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.GetAgentInfoDetailsFailed, 0)
		return nil
	}

	var portStatus string
	if req.PortStatus == "on" {
		portStatus = "LISTEN"
	} else if req.PortStatus == "off" {
		portStatus = ""
	}

	for _, portInfo := range agentInfoDetails.PortInfos {
		if (req.PortType == "" || req.PortType == portInfo.PortType) && (req.PortStatus == "" || portStatus == portInfo.PortStatus) {
			resp.Results = append(resp.Results, formjson.AgentPortInfoItem{
				Port:        portInfo.Port,
				PortType:    portInfo.PortType,
				PortStatus:  portInfo.PortStatus,
				PortService: portInfo.PortService,
				Pid:         portInfo.PortProcessInfo.Pid,
				PortProcessInfo: formjson.PortProcessInfo{
					Pid:           portInfo.PortProcessInfo.Pid,
					Cmdline:       portInfo.PortProcessInfo.Cmdline,
					CpuPercent:    webutils.Decimal(portInfo.PortProcessInfo.CpuPercent),
					MemoryPercent: webutils.Decimal(float64(portInfo.PortProcessInfo.MemoryPercent)),
					CreateTime:    portInfo.PortProcessInfo.CreateTime / 1000,
					Cwd:           portInfo.PortProcessInfo.Cwd,
					Username:      portInfo.PortProcessInfo.Username,
				},
			})
		}
	}

	support.SendApiResponse(ctx, resp, "")
	return
}

// DeleteAgentHandler
func DeleteAgentHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.DeleteAgentReq)
	resp := formjson.StatusResp{Status: "OK"}
	query := bson.M{
		"hash_id": bson.M{
			"$in": req.HashIds,
		},
	}
	if err = mongo.Agent.Delete(ctx, query); err != nil {
		mlog.Error("delete agent info failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.DeleteAgentInfoFailed, 0)
		return nil
	}
	support.SendApiResponse(ctx, resp, "")
	return
}

// AddAgentSecGrpRule
func AddAgentSecGrpRuleHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.AddAgentSecGrpRuleReq)
	resp := formjson.StatusResp{Status: "OK"}

	//数据校验
	port_int, _ := strconv.Atoi(req.Port)
	if port_int < 0 || port_int > 65535 {
		support.SendApiErrorResponse(ctx, support.PortRangeOverflow, 0)
		return nil
	}
	if matched, _ := regexp.MatchString("^(?:(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\/([0-9]|[1-2]\\d|3[0-2])$", req.Cidr); !matched {
		support.SendApiErrorResponse(ctx, support.CIDRInvalid, 0)
		return nil
	}

	ruleId := webutils.String.Hash(
		utils.IntToString(req.Direction),
		utils.IntToString(req.ProtocolType),
		utils.IntToString(req.ApplyPolicy),
		req.Cidr,
		req.Port,
	)
	query := bson.M{
		"agent_id": req.AgentId,
		"rule_id":  ruleId,
	}

	if exist := mongo.SecGrp.Find(ctx, query, nil); exist {
		support.SendApiErrorResponse(ctx, support.AddRepeatSecGrpRule, 0)
		return nil
	}

	secGrpRuleDoc := forward.ForwardAgentActionReq{
		Action:        forward.Action_ADDRULE,
		RuleDirection: int32(req.Direction),
		ChannelId:     req.AgentId,
		Cidr:          req.Cidr,
		ProtocolType:  int32(req.ProtocolType),
		Port:          req.Port,
		ApplyPolicy:   int32(req.ApplyPolicy),
	}
	if _, err = rpc.SendSecGrpRule(ctx.Request().Context(), &secGrpRuleDoc); err != nil {
		mlog.Error("rpc send sec grp rule failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.AddSecGrpRuleFailed, 0)
		return nil
	}

	secGrpDoc := models.SecGrp{
		AgentId:      req.AgentId,
		RuleId:       ruleId,
		Direction:    req.Direction,
		ProtocolType: req.ProtocolType,
		Port:         port_int,
		Cidr:         req.Cidr,
		ApplyPolicy:  req.ApplyPolicy,
		CreateTime:   time.Now().Unix(),
	}
	if err = mongo.SecGrp.Insert(ctx, &secGrpDoc); err != nil {
		mlog.Error("secgrp rule insert failed", zap.Error(err))
	}

	support.SendApiResponse(ctx, resp, "")
	return nil
}

// DeleteAgentSecGrpRuleHandler
func DeleteAgentSecGrpRuleHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.DeleteAgentSecGrpRuleReq)
	resp := formjson.StatusResp{Status: "OK"}
	//判断规则是否存在
	query := bson.M{
		"agent_id": req.AgentId,
		"rule_id":  req.RuleId,
	}

	var secGrpDoc models.SecGrp
	if exist := mongo.SecGrp.Find(ctx, query, &secGrpDoc); !exist {
		support.SendApiErrorResponse(ctx, support.SecGrpRuleNotExist, 0)
		return nil
	}

	//请求rpc
	secGrpRuleDoc := forward.ForwardAgentActionReq{
		Action:        forward.Action_DELETERULE,
		RuleDirection: int32(secGrpDoc.Direction),
		ChannelId:     req.AgentId,
		Cidr:          secGrpDoc.Cidr,
		ProtocolType:  int32(secGrpDoc.ProtocolType),
		Port:          utils.IntToString(secGrpDoc.Port),
		ApplyPolicy:   int32(secGrpDoc.ApplyPolicy),
	}
	if _, err = rpc.SendSecGrpRule(ctx.Request().Context(), &secGrpRuleDoc); err != nil {
		mlog.Error("rpc send sec grp rule failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.DeleteSecGrpRuleFailed, 0)
		return nil
	}
	//更新规则数据
	if err = mongo.SecGrp.Delete(ctx, query); err != nil {
		mlog.Error("secgrp rule delete failed", zap.Error(err))
	}

	support.SendApiResponse(ctx, resp, "")
	return
}

// ListAgentSecGrpHandler
func ListAgentSecGrpHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.ListAgentSecGrpReq)
	resp := formjson.ListAgentSecGrpResp{}

	query := bson.M{
		"agent_id": req.AgentId,
	}
	if req.Direction != 0 {
		query["direction"] = req.Direction
	}
	sortLists := make([]string, 0)
	if req.Sort != "" {
		sortLists = append(sortLists, req.Sort)
	} else {
		sortLists = append(sortLists, "-create_time")
	}

	var secGrpDocs []models.SecGrp
	if resp.Count, err = mongo.SecGrp.List(ctx, query, &secGrpDocs, req.Page, req.PageSize, sortLists...); err != nil {
		mlog.Error("list sec grp rule failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ListSecGrpRuleFailed, 0)
		return nil
	}

	for _, secGrpRule := range secGrpDocs {
		resp.Results = append(resp.Results, formjson.ListAgentSecGrpItem{
			RuleId:     secGrpRule.RuleId,
			Cidr:       secGrpRule.Cidr,
			CreateTime: secGrpRule.CreateTime,
			Port:       secGrpRule.Port,
			Direction: func(i int) (o string) {
				switch i {
				case -1:
					return support.SecGrp_RULE_IN
				case 1:
					return support.SecGrp_RULE_OUT
				default:
				}
				return support.SecGrp_Unknow
			}(secGrpRule.Direction),
			ProtocolType: func(i int) (o string) {
				switch i {
				case -1:
					return "ICMP"
				case 0:
					return "UDP"
				case 1:
					return "TCP"
				default:
				}
				return support.SecGrp_Unknow
			}(secGrpRule.ProtocolType),
			ApplyPolicy: func(i int) (o string) {
				switch i {
				case -1:
					return support.SecGrp_Policy_DROP
				case 0:
					return support.SecGrp_Policy_REJECT
				case 1:
					return support.SecGrp_Policy_ACCEPT
				default:
				}
				return support.SecGrp_Unknow
			}(secGrpRule.ApplyPolicy),
		})
	}

	support.SendApiResponse(ctx, resp, "")
	return
}

func StartBaselineScanHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.StartBaselineScanReq)
	resp := formjson.StatusResp{Status: "OK"}
	baselineScanDoc := forward.ForwardAgentActionReq{
		Action:    forward.Action_STARTCISCHECK,
		ChannelId: req.AgentId,
	}
	//发送扫描请求
	if _, err = rpc.SendSecGrpRule(ctx.Request().Context(), &baselineScanDoc); err != nil {
		mlog.Error("rpc send baseline scan action failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.StartBaselineScanFailed, 0)
		return nil
	}
	//开启同步等待并获取后台状态
	if isSuccess := sync.OpenWait(cache.BaselineSync(req.AgentId)); !isSuccess {
		support.SendApiErrorResponse(ctx, support.BaselineScanFailed, 0)
		return nil
	}
	support.SendApiResponse(ctx, resp, "")
	return
}

func ListAgentBaselineHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.ListAgentBaselineReq)

	var scanOutlineDoc models.CisScanOutline
	id := mongo.Baseline.FindCount(ctx, req.AgentId)
	if _, err = mongo.Baseline.Find(ctx, req.AgentId, id, &scanOutlineDoc); err != nil {
		mlog.Error("get baseline scan outline failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.GetBaselineScanResultFailed, 0)
		return nil
	}
	var scanResultDoc []models.CisScanResultItem
	query := bson.M{}
	if req.Status != "" {
		if req.Status == "1" {
			query["status"] = true
		} else {
			query["status"] = false
		}
	}
	query["agent_id"] = req.AgentId
	query["id"] = id
	var count int
	if count, err = mongo.Baseline.FindResult(ctx, query, req.Page, req.PageSize, &scanResultDoc); err != nil {
		mlog.Error("get baseline scan results failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.GetBaselineScanResultFailed, 0)
		return nil
	}
	resp := formjson.ListAgentBaselineResp{
		StartTime:    scanOutlineDoc.StartTime,
		EndTime:      scanOutlineDoc.EndTime,
		Count:        scanOutlineDoc.Count,
		SuccessCount: scanOutlineDoc.SuccessCount,
		FailedCount:  scanOutlineDoc.Count - scanOutlineDoc.SuccessCount,
		DisplayCount: count,
	}
	for _, item := range scanResultDoc {
		resp.Results = append(resp.Results, formjson.ListAgentBaselineRespItem{
			Id:     item.CisId,
			Status: item.Status,
			Desc: func(cisId string) string {
				var cisInfoDoc models.TbRepoCis
				if err := mongo.Baseline.FindRepoByCisId(ctx, cisId, &cisInfoDoc); err != nil {
					mlog.Error("find baseline repo info failed", zap.Error(err))
					return support.Unknow
				}
				return cisInfoDoc.Name
			}(item.CisId),
			IsIgnored: item.IsIgnored,
		})
	}

	support.SendApiResponse(ctx, resp, "")
	return
}

func UpdateAgentBaselineHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.UpdateAgentBaselineReq)
	resp := formjson.StatusResp{Status: "OK"}

	var scanOutlineDoc models.CisScanOutline
	id := mongo.Baseline.FindCount(ctx, req.AgentId)
	if exist, _ := mongo.Baseline.Find(ctx, req.AgentId, id, &scanOutlineDoc); !exist {
		mlog.Error("get baseline scan outline failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.GetBaselineScanResultFailed, 0)
		return nil
	}
	var scanResultDoc models.CisScanResultItem
	query := bson.M{
		"agent_id": req.AgentId,
		"cis_id":   req.CisId,
		"id":       id,
	}
	if exist, _ := mongo.Baseline.FindOneResult(ctx, query, &scanResultDoc); !exist {
		mlog.Error("get baseline scan result failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.GetBaselineScanResultFailed, 0)
		return nil
	}

	if scanResultDoc.Status {
		support.SendApiErrorResponse(ctx, support.UpdateBaselineInvalid, 0)
		return nil
	}

	updateOutline := bson.M{
		"success_count": func(status bool) int {
			if status {
				return scanOutlineDoc.SuccessCount - 1
			}
			return scanOutlineDoc.SuccessCount + 1
		}(scanResultDoc.IsIgnored),
	}
	updateResult := bson.M{
		"is_ignored": !scanResultDoc.IsIgnored,
	}
	if err = mongo.Baseline.UpdateOutline(ctx, req.AgentId, id, updateOutline); err != nil {
		mlog.Error("update baseline scan outline failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.UpdateBaselineStatusFailed, 0)
		return nil
	}
	if err = mongo.Baseline.UpdateResult(ctx, query, updateResult); err != nil {
		mlog.Error("update baseline scan result failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.UpdateBaselineStatusFailed, 0)
		return nil
	}

	support.SendApiResponse(ctx, resp, "")
	return
}

func GetBaselineInfoHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.GetBaselineInfoReq)

	var cisDoc models.TbRepoCis
	if err := mongo.Baseline.FindRepoByCisId(ctx, req.CisId, &cisDoc); err != nil {
		mlog.Error("get baseline info failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.GetBaselineInfoFailed, 0)
		return nil
	}

	resp := formjson.GetBaselineInfoResp{
		Name:    cisDoc.Name,
		Desc:    cisDoc.Description,
		Explain: cisDoc.Rationale,
		Solute:  cisDoc.Remediation,
	}

	support.SendApiResponse(ctx, resp, "")
	return
}
