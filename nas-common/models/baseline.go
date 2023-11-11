//---------------------------------
//File Name    : ../../nas-common/models/baseline.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2022-01-04 21:32:31
//Description  :
//----------------------------------
package models

type CisScanResultItem struct {
	Id        int    `bson:"id" mgo:"index:2"`       //id
	AgentId   string `bson:"agent_id" mgo:"index:1"` //agent_id
	CisId     string `bson:"cis_id"`                 //cis id
	Status    bool   `bson:"status"`                 //状态(是否符合cis rule)
	IsIgnored bool   `bson:"is_ignored"`             //是否会忽略
}

func (CisScanResultItem) Collection() string {
	return "tb_scan_baseline_results"
}

type CisScanOutline struct {
	AgentId      string `bson:"agent_id" mgo:"index:1"` //agent_id
	EndTime      int64  `bson:"end_time"`               //扫描结束时间
	StartTime    int64  `bson:"start_time"`             //扫描开始时间
	Count        int    `bson:"count"`                  //总数
	SuccessCount int    `bson:"success_count"`          //合规数量
	Id           int    `bson:"id" mongo:"index:2"`     //id
}

func (CisScanOutline) Collection() string {
	return "tb_scan_baseline"
}

type TbRepoCis struct {
	Id          string `yaml:"id" bson:"id" mgo:"index:1"`     //id
	Name        string `yaml:"name" bson:"name"`               //名字
	Description string `yaml:"description" bson:"description"` //描述
	Rationale   string `yaml:"rationale" bson:"rationale"`     //解释
	Remediation string `yaml:"remediation" bson:"remediation"` //建议
}

func (TbRepoCis) Collection() string {
	return "tb_repo_cis"
}
