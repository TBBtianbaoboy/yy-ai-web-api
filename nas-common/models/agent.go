//---------------------------------
//File Name    : agent.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-21 18:43:13
//Description  :
//----------------------------------
package models

type AgentHandlerInfo struct {
	HashId     string `bson:"hash_id" mgo:"index:1"` //agent 唯一ID MD5("agent":MacAddress)
	AgentIp    string `bson:"agent_ip"`              //agent IP
	Hostname   string `bson:"hostname"`              //agent hostname
	Pid        int64  `bson:"pid"`                   //agent pid
	MacAddress string `bson:"mac_address"`           //agent mac address
	Jointime   int64  `bson:"join_time"`             //agent 加入时间
	UpdateTime int64  `bson:"update_time"`           //更新时间
	IsDeleted  bool   `bson:"is_deleted"`            //软删除
}

func (a *AgentHandlerInfo) Collection() string {
	return "tb_agent"
}

//-------------------------------------------------------------

type AgentSystemResources struct {
	CpuUsed     float64 `bson:"cpu_used"`
	CpuCore     int32   `bson:"cpu_core"`
	MemoryUsed  float64 `bson:"memory_used"`
	MemoryTotal uint64  `bson:"memory_total"`
}

type PortProcessInfo struct {
	Pid           int64   `bson:"pid"`            //进程PID
	Cmdline       string  `bson:"cmdline"`        //进程运行命令
	CpuPercent    float64 `bson:"cpu_percent"`    //cpu使用时间百分比
	CreateTime    int64   `bson:"create_time"`    //创建时间
	Cwd           string  `bson:"cwd"`            //工作目录
	MemoryPercent float32 `bson:"memory_percent"` //内存使用百分比
	Username      string  `bson:"username"`       //所属用户
}

type AgentPortInfo struct {
	PortType        string          `bson:"port_type"`         // 端口类型
	Port            int64           `bson:"port"`              //端口号
	PortStatus      string          `bson:"port_status"`       //端口状态
	PortService     string          `bson:"port_service"`      //端口服务
	PortProcessInfo PortProcessInfo `bson:"port_process_info"` //端口进程信息
}

type AgentDiskInfo struct {
	Device      string  `bson:"device"`      //设备名
	MountPoint  string  `bson:"mountpoint"`  //挂载点
	Fstype      string  `bson:"fstype"`      //文件系统类型
	Options     string  `bson:"options"`     //其他信息
	Total       uint64  `bson:"total"`       //磁盘总量
	UsedPercent float64 `bson:"usedpercent"` //磁盘使用率
}

type AgentHandlerInfoDetails struct {
	HashId    string               `bson:"hash_id"`    //agent 唯一ID
	DiskInfo  AgentDiskInfo        `bson:"disk_info"`  //系统磁盘信息
	Resources AgentSystemResources `bson:"resources"`  //系统资源信息
	PortInfos []AgentPortInfo      `bson:"port_infos"` //agent 所在主机port信息
}

func (a *AgentHandlerInfoDetails) Collection() string {
	return "tb_agent_info"
}
