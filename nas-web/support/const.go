// ---------------------------------
// File Name    : const.go
// Author       : aico
// Mail         : 2237616014@qq.com
// Github       : https://github.com/TBBtianbaoboy
// Site         : https://www.lengyangyu520.cn
// Create Time  : 2021-12-14 14:49:24
// Description  :
// ----------------------------------
package support

// jwt
const (
	Auth           = "Nas-Auth"
	JWTFixedSecKey = "#J5iAAM*F!kY%~TH]DF_-C#jddXmWgf?"
)

// scan type
const (
	ScanType_FastMode    string = "快速扫描"
	ScanType_MostCommon  string = "普遍扫描"
	ScanType_PortRange   string = "端口范围扫描"
	ScanType_PortSingle  string = "单个端口扫描"
	ScanType_PortService string = "端口服务扫描"
	ScanType_Default     string = "默认扫描"

	SecGrp_RULE_IN       string = "入方向"
	SecGrp_RULE_OUT      string = "出方向"
	SecGrp_Policy_DROP   string = "丢弃"
	SecGrp_Policy_REJECT string = "拒绝"
	SecGrp_Policy_ACCEPT string = "允许"
	SecGrp_Unknow        string = "Error"

	ChatMessageRoleSystem    = "system"
	ChatMessageRoleUser      = "user"
	ChatMessageRoleAssistant = "assistant"
)
