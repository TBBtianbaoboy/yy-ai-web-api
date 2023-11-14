package formjson

type LoginReq struct {
	CaptId   string `json:"capt_id" validate:"required"`  //验证码唯一ID
	Username string `json:"username" validate:"required"` // 用户名称
	Password string `json:"password" validate:"required"` // 用户密码
	Vcode    string `json:"vcode" validate:"required"`    //验证码
}

type AddUserReq struct {
	Username string `json:"username"  validate:"required"` //用户名
	UserType int    `json:"user_type"`                     //用户类型
	Password string `json:"password"  validate:"required"` //用户密码
	Confirm  string `json:"confirm"  validate:"required"`  //确认用户密码
	Mail     string `json:"mail"`                          //用户邮箱
	Mobile   string `json:"mobile"`                        //用户手机号码
	Remark   string `json:"remark"`                        //用户备注信息
}

type DeleteUserReq struct {
	Uids []int `json:"uids" validate:"required"` //要删除的用户id列表
}

type EditUserReq struct {
	Uid    int    `json:"uid" validate:"required"` //用户ID
	Mail   string `json:"mail"`                    //用户邮箱
	Mobile string `json:"mobile"`                  //用户联系方式
}

type ListUserReq struct {
	Page     int    `json:"page" form:"page"`           //分页
	PageSize int    `json:"page_size" form:"page_size"` //分页大小
	Search   string `json:"search" form:"search"`       //搜索 [用户名| 邮箱|联系方式]
	Enable   int    `json:"enable" form:"enable"`       //是否允许登录 [1 允许登录| -1 禁止登录| 0 全部]
	UserType int    `json:"user_type" form:"user_type"` //用户类型 [ 1 管理员 | 2 超级用户 | 3 普通用户| 0 全部用户]
}

type UpdateUserPasswdReq struct {
	Uid  int    `json:"uid"`  //用户ID
	Old  string `json:"old"`  //旧密码
	New  string `json:"new"`  //新密码
	New2 string `json:"new2"` //确认新密码
}

type UpdateUserStatusReq struct {
	Uid    int  `json:"uid" validate:"required"` //用户id
	Enable bool `json:"enable"`                  //是否允许登录
}

// audio [Ok]
type TranscriptionsReq struct {
	Language string `json:"language" form:"language"` // input audio language [optional]
}

// audio [Ok]
type SpeechReq struct {
	ModelName string  `json:"model_name"` // audio model name < tts-1 | tts-1-hd, {optional,default tts-1} >
	Input     string  `json:"input"`      // input text used to generate audio == < user input, {required,not empty} >
	Voice     string  `json:"voice"`      // voice type of generated audio == < alloy | echo | fable | onyx | nova | shimmer, {required,not empty} >
	Speed     float64 `json:"speed"`      // speed of generated audio == < 0.25-4.0 {optional,default 1.0} >
	Format    string  `json:"format"`     // format of generated audio == < mp3 | opus | aac | flac, {optional,default mp3} >
}

// chat [Ok]
type SendNoContextNoStreamChatReq struct {
	ModelName string `json:"model_name"` // chat model name
	// System    string `json:"system"`     // model system content
	Question string `json:"question"` // chat content
}

// chat [Ok]
type SendNoContextStreamChatReq struct {
	ModelName string `json:"model_name"` // chat model name
	Question  string `json:"question"`   // chat content
}

// chat [Ok]
type SendContextStreamChatReq struct {
	ModelName string `json:"model_name"` // chat model name
	Question  string `json:"question"`   // chat content
	SessionId int    `json:"session_id"` // session id
}

// chat [Ok]
type DeleteContextStreamChatReq struct {
	SessionId int `json:"session_id"` // session id
}

type DeleteScanReq struct {
	ScanIds []string `json:"scan_ids"` //to be deleted scan id set
}

type ListScanReq struct {
	Page     int    `json:"page" form:"page"`           //分页
	PageSize int    `json:"page_size" form:"page_size"` // 分页大小
	Sort     string `json:"sort" form:"sort"`           //按照开始扫描的时间进行排序(-start_time|start_time)
	Search   string `json:"search" form:"search"`       //按照扫描域进行搜索
}

type GetFirstScanResultReq struct {
	ScanId string `form:"scan_id" json:"scan_id"` //scan id
}

type GetScanIpResultReq struct {
	ScanId string `form:"scan_id" json:"scan_id"`
	Ip     string `form:"ip" json:"ip"`
}

type ListAgentReq struct {
	Page     int    `form:"page" json:"page"`           //分页
	PageSize int    `form:"page_size" json:"page_size"` //分页大小
	Search   string `form:"search" json:"search"`       //根据主机IP/hostname 进行搜索
}

// image [Ok]
type GenerateImageReq struct {
	ModelName string `json:"model_name"` // model name
	Prompt    string `json:"prompt"`     // prompt content
	Size      string `json:"size"`       //image size
	Quality   string `json:"quality"`    // image quality
}

type AgentSystemInfoReq struct {
	HashId string `form:"hash_id" json:"hash_id" validate:"required"` //agent hash id (唯一)
}

type AgentPortInfoReq struct {
	HashId     string `form:"hash_id" json:"hash_id" validate:"required"` //agent hash id (唯一)
	PortType   string `form:"port_type" json:"port_type"`                 //port type
	PortStatus string `form:"port_status" json:"port_status"`             //port_status
	Search     string `form:"search" json:"search"`                       //search
}

type DeleteAgentReq struct {
	HashIds []string `json:"hash_ids"` //agent hash id (唯一)
}

type AddAgentSecGrpRuleReq struct {
	Direction    int    `json:"direction" form:"direction"`         //方向 (-1 入方向 | 1 出方向)
	ProtocolType int    `json:"protocol_type" form:"protocol_type"` //协议类型 (1 tcp | 0 udp | -1 icmp)
	Port         string `json:"port" form:"port"`                   //端口(0-65535)
	ApplyPolicy  int    `json:"apply_policy" form:"apply_policy"`   //授权策略 (1 ACCEPT | 0 REHECT | DROP)
	Cidr         string `json:"cidr" form:"cidr"`                   //授权对象 (such as 0.0.0.0/24)
	AgentId      string `json:"agent_id" form:"agent_id"`           //agent id
}

type DeleteAgentSecGrpRuleReq struct {
	AgentId string `json:"agent_id" form:"agent_id"` //agent id
	RuleId  string `json:"rule_id" form:"rule_id"`   //rule id
}

type ListAgentSecGrpReq struct {
	AgentId   string `json:"agent_id" form:"agent_id"`   //agent_id
	Page      int    `json:"page" form:"page"`           //分页
	PageSize  int    `json:"page_size" form:"page_size"` //分页大小
	Direction int    `json:"direction" form:"direction"` //根据控制方向来获取数据(-1 入|0 全部| 1 出)
	Sort      string `json:"sort" form:"sort"`           //根据时间|端口号进行排序
}

type StartBaselineScanReq struct {
	AgentId string `json:"agent_id" form:"agent_id"` //agent_id
}

type ListAgentBaselineReq struct {
	Page     int    `json:"page" form:"page"`                             //分页
	PageSize int    `json:"page_size" form:"page_size"`                   //分页大小
	AgentId  string `json:"agent_id" form:"agent_id" validate:"required"` //agent_id
	Status   string `json:"status" form:"status"`                         //基线扫描状态("0" failed | "1" success)
}

type UpdateAgentBaselineReq struct {
	AgentId string `json:"agent_id" form:"agent_id" validate:"required"` //agent_id
	CisId   string `json:"cis_id" form:"cis_id" validate:"required"`     //cis_id
}

type GetBaselineInfoReq struct {
	CisId string `json:"cis_id" form:"cis_id" validate:"required"` //cis_id
}

type ResetPasswdReq struct {
	Uid      int    `json:"uid" form:"uid" validate:"required"`           //用户ID
	Password string `json:"password" form:"password" validate:"required"` //重置用户密码
}
