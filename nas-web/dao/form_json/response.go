// ---------------------------------
// File Name    : response.go
// Author       : aico
// Mail         : 2237616014@qq.com
// Github       : https://github.com/TBBtianbaoboy
// Site         : https://www.lengyangyu520.cn
// Create Time  : 2021-12-14 18:46:30
// Description  :
// ----------------------------------
package formjson

import "nas-common/models"

type StatusResp struct {
	Status string `json:"status"` //基本状态
}

// service auth.go
type VerifyCodeResp struct {
	CaptId string `json:"capt_id"` //验证码唯一ID
	Image  string `json:"image"`   //验证码图片数据
}

// user.go [Ok]
type LoginResp struct {
	Uid           int    `json:"uid"`           //用户Id
	Enable        bool   `json:"enable"`        //是否允许登陆
	Authorization string `json:"authorization"` //用于鉴权的token
	Username      string `json:"username"`      //用户名
}

// service scan.go
type ListScanItem struct {
	ScanId            string `bson:"scan_id" json:"scan_id"`                           //扫描项唯一ID
	StartTime         int64  `bson:"start_time" json:"start_time"`                     // 开始扫描时间
	ScanIp            string `bson:"scan_ip" json:"scan_ip"`                           // 扫描域
	ScanType          string `bson:"scan_type" json:"scan_type"`                       //扫描类型
	IsFirstScanDone   bool   `bson:"is_first_scan_done" json:"is_first_scan_done"`     // 预扫是否结束
	FirstScanDoneTime int64  `bson:"first_scan_done_time" json:"first_scan_done_time"` // 预扫结束时间
	IsDeepScanDone    bool   `bson:"is_deep_scan_done" json:"is_deep_scan_done"`       // 深扫是否结束
	DeepScanDoneTime  int64  `bson:"deep_scan_done_time" json:"deep_scan_done_time"`   // 深扫结束时间
	Status            int    `bson:"status" json:"status"`                             //扫描结果状态 -1 表示扫描失败 | 0 表示正在扫描中 | 1 表示扫描成功
}

// service scan.go
type ListScanResp struct {
	Count   int            `bson:"count" json:"count"`     //总数
	Results []ListScanItem `bson:"results" json:"results"` //扫描列表
}

// chat [Ok]
type SendNoContextNoStreamChatResp struct {
	Answer string `json:"answer"` //assistant answer
}

// chat [Ok]
type SendNoContextStreamChatResp struct {
	Answer string `json:"answer"` //assistant answer
}

// service scan.go
type GetFirstScanResultResp struct {
	All             string           `json:"all"`       //all count
	Rate            int              `json:"rate"`      //success rate
	Count           string           `json:"count"`     //been scaned ip nums
	IpDomain        string           `json:"ip_domain"` //ip domain
	Elapsed         float32          `json:"elapsed"`   // used time
	PortFB          map[uint16]int   `json:"portfb"`
	ServiceFB       map[string]int   `json:"servicefb"`
	StartTime       int64            `json:"start_time"`
	EndTime         int64            `json:"end_time"`
	Results         []DeepScanResult `json:"results"`
	ScanType        int32            `json:"scan_type"`
	ScanTypeMessage string           `json:"scan_type_message"`
	WithOs          bool             `json:"with_os"`
	WithScript      bool             `json:"with_script"`
	WithTrace       bool             `json:"with_trace"`
	WithService     bool             `json:"with_service"`
}

// service scan.go
type GetScanIpResultResp struct {
	PortService    []PortServiceInfo `json:"port_service_info"`
	PortServiceSum int               `json:"port_service_sum"`
	OsGuest        []OsGuestInfo     `json:"os_guest"`
	OsGuestSum     int               `json:"os_guest_sum"`
	Trace          []string          `json:"trace"`
	TraceTTL       []float32         `json:"trace_ttl"`
	TracePort      string            `json:"trace_port"`
	TraceProtocol  string            `json:"trace_protocol"`
}

type ScriptInfo struct {
	Id       string           `json:"id"`
	Output   string           `json:"output"`
	Elements []models.Element `json:"elements"`
	Tabels   []models.Table   `json:"tables"`
}

type OsGuestInfo struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type PortServiceInfo struct {
	PortId           uint16           `json:"port_id"`
	Protocol         string           `json:"protocol"`
	ServiceName      string           `json:"service_name"`
	ServiceExtraInfo PortServiceRxtra `json:"service_extra_info"`
	ServiceState     string           `json:"service_state"`
	ServiceTTL       float32          `json:"service_ttl"`
	Script           []ScriptInfo     `json:"script"`
}

type PortServiceRxtra struct {
	Version   string   `json:"version"`
	Product   string   `json:"product"`
	Method    string   `json:"method"`
	ExtraInfo string   `json:"extra_info"`
	Tunnel    string   `json:"tunnel"`
	Cpes      []string `json:"cpes"`
}

type DeepScanResult struct {
	Ip           string `json:"ip"`
	IpType       string `json:"ip_type"`
	Status       string `json:"status"`
	StatusReason string `json:"status_reason"`
}

// service agent.go
type ListAgentItem struct {
	HashId     string `json:"hash_id"`     //agent hash id (唯一)
	AgentIp    string `json:"agent_ip"`    //agent ip
	Hostname   string `json:"hostname"`    //agent hostname
	Pid        int64  `json:"pid"`         //agent pid
	JoinTime   int64  `json:"join_time"`   //agent 加入时间
	UpdateTime int64  `json:"update_time"` //agent 更新时间
	Status     bool   `json:"status"`      //agent 状态
}

type ListAgentResp struct {
	Count   int             `json:"count"`   //总数
	Results []ListAgentItem `json:"results"` //结果
}

// service agent.go
type AgentSystemInfoDisk struct {
	Device      string  `json:"device"`       //设备名
	MountPoint  string  `json:"mount_point"`  //挂载点
	Fstype      string  `json:"fstype"`       //文件系统类型
	Options     string  `json:"options"`      //其他信息
	Total       string  `json:"total"`        //磁盘总量
	UsedPercent float64 `json:"used_percent"` //磁盘使用率
}

type AgentSystemInfoResp struct {
	CpuUsed     float64             `json:"cpu_used"`     //cpu 使用率
	MemoryUsed  float64             `json:"memory_used"`  //内存使用率
	CpuCore     int32               `json:"cpu_core"`     //cpu 核数
	MemoryTotal string              `json:"memory_total"` //内存总量
	DiskInfo    AgentSystemInfoDisk `json:"disk_info"`    //磁盘信息
}

// service agent.go
type PortProcessInfo struct {
	Pid           int64   `json:"pid_num"`        //PID
	Cmdline       string  `json:"cmdline"`        //启动进程的命令
	CpuPercent    float64 `json:"cpu_percent"`    //cpu使用时间百分比
	CreateTime    int64   `json:"create_time"`    //创建时间
	Cwd           string  `json:"cwd"`            //工作目录
	MemoryPercent float64 `json:"memory_percent"` //内存使用百分比
	Username      string  `json:"username"`       //所属用户
}

type AgentPortInfoItem struct {
	Port            int64           `json:"port"`              //端口
	PortType        string          `json:"port_type"`         //端口类型
	PortService     string          `json:"port_service"`      //端口服务
	PortStatus      string          `json:"port_status"`       //端口状态
	Pid             int64           `json:"pid"`               //进程PID
	PortProcessInfo PortProcessInfo `json:"port_process_info"` //与端口对应的进程信息
}

type AgentPortInfoResp struct {
	Results []AgentPortInfoItem `json:"results"` //结果
}

type DownloadAgentResp struct {
	Status string `json:"status"` //状态
}

// service user.go
type ListUserItem struct {
	Uid        int    `json:"uid"`         //用户id
	Username   string `json:"username"`    //用户名
	Enable     bool   `json:"enable"`      //是否允许登录
	UserType   int    `json:"user_type"`   //用户类型
	Mail       string `json:"mail"`        //用户邮箱
	Mobile     string `json:"mobile"`      //用户联系方式
	CreateTime int64  `json:"create_time"` //用h用户创建时间
}

type ListUserResp struct {
	Count   int            `json:"count"`   //总数
	Results []ListUserItem `json:"results"` //结果
}

type GetUserInfoResp struct {
	Uid        int    `json:"uid" form:"uid"`                 //用户id
	UserType   int    `json:"user_type" form:"user_type"`     //用户角色(1 管理员|2 超级用户|3 普通用户)
	UserName   string `json:"username" form:"username"`       //用户名
	Email      string `json:"email" form:"email"`             //注册邮箱
	Phone      string `json:"phone" form:"phone"`             //联系手机
	PS         int    `json:"ps" form:"ps"`                   //密码强度(1 低|2 中|3 强)
	CreateTime int64  `json:"create_time" form:"create_time"` //创建时间
}

// service agent.go
type ListAgentSecGrpItem struct {
	RuleId       string `json:"rule_id" form:"rule_id"`             //唯一表示此条规则 rule_id
	Direction    string `json:"direction" form:"direction"`         //规则作用方向 (-1 In | 1 Out)
	ProtocolType string `json:"protocol_type" form:"protocol_type"` //协议类型
	ApplyPolicy  string `json:"apply_policy" form:"apply_policy"`   //授权策略
	Port         int    `json:"port" form:"port"`                   //作用端口
	Cidr         string `json:"cidr" form:"cidr"`                   //授权对象
	CreateTime   int64  `json:"create_time" form:"create_time"`     //创建时间
}

type ListAgentSecGrpResp struct {
	Count   int                   `json:"count" form:"count"`     //数量
	Results []ListAgentSecGrpItem `json:"results" form:"results"` //结果
}

// service agent.go
type ListAgentBaselineResp struct {
	StartTime    int64                       `json:"start_time" form:"start_time"`       //扫描开始时间
	EndTime      int64                       `json:"end_time" form:"end_time"`           //扫描结束时间
	Count        int                         `json:"count" form:"count"`                 //检查项总数
	SuccessCount int                         `json:"success_count" form:"success_count"` //通过项数
	FailedCount  int                         `json:"failed_count" form:"failed_count"`   //失败项数
	Results      []ListAgentBaselineRespItem `json:"results" form:"results"`             //结果
	DisplayCount int                         `json:"display_count" form:"display_count"` //当前类别总数
}

type ListAgentBaselineRespItem struct {
	Id        string `json:"id" form:"id"`                 //检查项id
	Status    bool   `json:"status" form:"status"`         //是否合规
	Desc      string `json:"desc" form:"desc"`             //描述
	IsIgnored bool   `json:"is_ignored" form:"is_ignored"` //是否被忽略
}

// service agent.go
type GetBaselineInfoResp struct {
	Name    string `json:"name" form:"name"`       //基线规则名称
	Desc    string `json:"desc" form:"desc"`       //基线规则描述
	Explain string `json:"explain" form:"explain"` //解释
	Solute  string `json:"solute" form:"solute"`   //解决方案
}
