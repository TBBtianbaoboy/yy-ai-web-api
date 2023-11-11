//---------------------------------
//File Name    : scan_ip.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-10 13:29:09
//Description  :
//----------------------------------
package models

import (
	"time"
)

type ScanPortType struct {
	IsFastMode   bool   `json:"is_fast_mode"`
	IsMostCommon bool   `json:"is_most_common"`
	PortRange    string `json:"post_range"`
	PortSingle   string `json:"port_single"`
	PortService  string `json:"port_service"`
}

type ScanOptions struct {
	Type                ScanPortType `json:"type"`
	IsWithServiceInfo   bool         `json:"is_with_service_info"`
	IsWithOsDetection   bool         `json:"is_with_os_detection"`
	IsWithDefaultScript bool         `json:"is_with_default_script"`
	IsWithTraceRoute    bool         `json:"is_with_trace_route"`
}

// use choose scan info
type ScanIpInfo struct {
	IpId     string      `json:"ip_id"` // 唯一Id,全链路一致，产生于此
	IpDomain string      `json:"ip_domain"`
	Options  ScanOptions `json:"options"`
}

//--------------------------------------------------------

type DeepScanOptions struct {
	ScanType        int32  `bson:"scan_type"` // (-1 use default)
	ScanTypeMessage string `bson:"scan_type_message"`
	WithOs          bool   `bson:"with_os"`
	WithTrace       bool   `bson:"with_trace"`
	WithScript      bool   `bson:"with_script"`
	WithService     bool   `bson:"with_service"`
}

// first scan result
type FirstScanIpResult struct {
	ScanIpId       string          `bson:"_id"`              // 唯一id，全链路一致
	StartTime      time.Time       `bson:"start_time"`       // zmap 开始扫描时间
	IpDomain       string          `bson:"ip_domain"`        // zmap 扫描目标
	EndTime        time.Time       `bson:"end_time"`         // zmap 扫描结束时间
	Count          int             `bson:"count"`            // zmap 扫描结果数量
	Result         []string        `bson:"result"`           // zmap 扫描结果
	DeepScanOption DeepScanOptions `bson:"deep_scan_option"` // 深入扫描参数
}

func (FirstScanIpResult) Collection() string {
	return "tb_first_scan_result"
}

//---------------------------------------------------------

type Distance struct {
	Value int `bson:"value"`
}

type OS struct {
	PortsUsed []PortUsed `bson:"port_user"`
	Matches   []OSMatch  `bson:"matches"`
}

type PortUsed struct {
	State string `bson:"state"`
	Proto string `bson:"proto"`
	ID    int    `bson:"id"`
}

type OSMatch struct {
	Name     string    `bson:"name"`
	Accuracy int       `bson:"accuracy"`
	Classes  []OSClass `bson:"osclass"`
}

type OSClass struct {
	Vendor       string   `bson:"vendor"`
	OSGeneration string   `bson:"osgen"`
	Type         string   `bson:"type"`
	Accuracy     int      `bson:"accuracy"`
	Family       string   `bson:"osfamily"`
	CPEs         []string `bson:"cpe"`
}

type Status struct {
	State     string  `bson:"state"`
	Reason    string  `bson:"reason"`
	ReasonTTL float32 `bson:"reason_ttl"`
}

type Trace struct {
	Proto string `bson:"proto"`
	Port  int    `bson:"port"`
	Hops  []Hop  `bson:"hop"`
}

type Hop struct {
	TTL    float32 `bson:"ttl"`
	RTT    string  `bson:"rtt"`
	IPAddr string  `bson:"ipaddr"`
}

type Uptime struct {
	Seconds  int    `bson:"seconds"`
	Lastboot string `bson:"lastboot"`
}

type Address struct {
	Addr     string `bson:"addr"`
	AddrType string `bson:"addrtype"`
	Vendor   string `bson:"vendor"`
}

type ExtraPort struct {
	State   string   `bson:"state"`
	Count   int      `bson:"count"`
	Reasons []Reason `bson:"extrareasons"`
}

type Reason struct {
	Reason string `bson:"reason"`
	Count  int    `bson:"count"`
}

type Hostname struct {
	Name string `bson:"name"`
	Type string `bson:"type"`
}

type Port struct {
	ID       uint16   `bson:"portid"`
	Protocol string   `bson:"protocol"`
	Service  Service  `bson:"service"`
	State    State    `bson:"state"`
	Scripts  []Script `bson:"script"`
}

type Service struct {
	ExtraInfo string   `bson:"extrainfo"`
	Method    string   `bson:"method"`
	Name      string   `bson:"name"`
	Product   string   `bson:"product"`
	ServiceFP string   `bson:"servicefp"`
	Tunnel    string   `bson:"tunnel"`
	Version   string   `bson:"version"`
	CPEs      []string `bson:"cpe"`
}

type State struct {
	State     string  `bson:"state"`
	Reason    string  `bson:"reason"`
	ReasonTTL float32 `bson:"reason_ttl"`
}

type Script struct {
	ID       string    `bson:"id"`
	Output   string    `bson:"output"`
	Elements []Element `bson:"elem"`
	Tables   []Table   `bson:"table"`
}

type Table struct {
	Key      string    `bson:"key"`
	Tables   []Table   `bson:"table"`
	Elements []Element `bson:"elem"`
}

type Element struct {
	Key   string `bson:"key" json:"key"`
	Value string `bson:"value" json:"value"`
}

type Host struct {
	StartTime  string      `bson:"starttime"`  // 扫描单个目标主机的开始时间
	EndTime    string      `bson:"endtime"`    // 扫描单个目标主机的结束时间
	OS         OS          `bson:"os"`         // 开启操作系统检测所扫描到的数据 -- 开启才有
	Status     Status      `bson:"status"`     // 目标主机的状态
	Trace      Trace       `bson:"trace"`      // 开启路由追踪所扫描到的数据 -- 开启才有
	Uptime     Uptime      `bson:"uptime"`     // 目标主机的在线信息
	Addresses  []Address   `bson:"address"`    // 目标主机的地址信息
	ExtraPorts []ExtraPort `bson:"extraports"` // 目标主机的额外端口信息(close/filtered)
	Hostnames  []Hostname  `bson:"hostnames"`  // 目标主机的主机名信息
	Distance   Distance    `bson:"distance"`   // 到达目标主机的跳数 -- 开启路由追踪才有
	Ports      []Port      `bson:"ports"`      // 目标主机的端口信息
}

// deep scan result
type DeepScanIpResult struct {
	DeepScanId string    `bson:"_id"`          // 唯一ID，全链路保持一致
	DeepScanIp string    `bson:"deep_scan_ip"` // 扫描域
	Elapsed    float32   `bson:"elapsed"`      // 扫描耗时
	StartTime  time.Time `bson:"start_time"`   // 开始扫描的时间
	EndTime    time.Time `bson:"end_time"`     // 所有操作完成的结束时间
	Count      int       `bson:"count"`        // 扫描域中被扫描的IP 数- 小于或等于zmap 扫描到的IP
	Hosts      []Host    `bson:"hosts"`        // 扫描结果
}

func (DeepScanIpResult) Collection() string {
	return "tb_deep_scan_result"
}

//------------------------------------------------------

type ScanStatusDesc struct {
	ScanId            string    `bson:"_id"` // 唯一ID，全链路保持一致
	Uid               int       `bson:"uid"`
	StartTime         time.Time `bson:"start_time"`           // 开始扫描时间
	ScanIp            string    `bson:"scan_ip"`              // 扫描域
	ScanType          string    `bson:"scan_type"`            //扫描类型
	IsFirstScanDone   bool      `bson:"is_first_scan_done"`   // 预扫是否结束
	FirstScanDoneTime time.Time `bson:"first_scan_done_time"` // 预扫结束时间
	IsDeepScanDone    bool      `bson:"is_deep_scan_done"`    // 深扫是否结束
	DeepScanDoneTime  time.Time `bson:"deep_scan_done_time"`  // 深扫结束时间
	Status            int       `bson:"status"`               //扫描结果状态 -1 failure | 0 in scanning | 1 success
}

func (ScanStatusDesc) Collection() string {
	return "tb_scan_status_desc"
}
