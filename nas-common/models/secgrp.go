//---------------------------------
//File Name    : secgrp.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-30 12:49:50
//Description  :
//----------------------------------
package models

type SecGrpRule struct {
	Direction    int32  `json:"direction"`     // -1 (input) | 1 (output)
	ProtocolType int32  `json:"protocol_type"` //协议类型
	ApplyPolicy  int32  `json:"apply_policy"`  //授权策略 ("accept" | "reject" | "drop")
	Port         string `json:"port"`          //端口
	Cidr         string `json:"cidr"`          //授权对象
	Action       int32  `json:"action"`        //行为
}

type SecGrp struct {
	AgentId      string `bson:"agent_id" mgo:"index:1"`
	RuleId       string `bson:"rule_id" mgo:"index:2"` //唯一表示此条规则
	Direction    int    `bson:"direction"`             // (-1 入| 1 出)
	ProtocolType int    `bson:"protocol_type"`         // (1 tcp | 0 udp | -1 icmp)
	ApplyPolicy  int    `bson:"apply_policy"`          // (1 ACCEPT | 0 REJECT | -1 DROP)
	Port         int    `bson:"port"`
	Cidr         string `bson:"cidr"`
	CreateTime   int64  `bson:"create_time"` //创建时间
}

func (s *SecGrp) Collection() string {
	return "tb_secgrp"
}
