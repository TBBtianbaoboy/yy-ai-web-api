package support

const (
	VCodeFailed                string = "验证码错误"
	TokenParseFailedAndEmpty   string = "解析错误,token为空"
	TokenExpire                string = "用户会话认证失败"
	TokenFlushFailed           string = "token续期失败"
	TokenParseFailedAndInvalid string = "解析错误,token无效"

	UserNameFailed     string = "用户名错误"
	PasswordFailed     string = "密码错误"
	PasswordDiffFailed string = "两次密码不相同"
	UserLockFailed     string = "账户已锁定"
	CanNotLogin        string = "禁止登录"

	NoPrivilegeAddUser     string = "无创建用户权限"
	PasswordNotConfirm     string = "两次密码不匹配，请确认"
	UserIsExist            string = "同名用户已存在"
	UserNotExist           string = "用户不存在"
	PasswordStrengthFailed string = "密码强度过低"
	CreateUserFailed       string = "用户创建失败"
	DeleteUsersFailed      string = "删除用户失败"
	EditUserFailed         string = "编辑用户失败"
	ListUsersFailed        string = "获取用户列表失败"
	UpdateUserPasswdFailed string = "修改用户密码失败"
	UpdateUserStatusFailed string = "修改用户状态失败"

	GetUserInfoFailed   string = "获取用户信息失败"
	EditUserInfoFailed  string = "编辑用户信息失败"
	GenerateTokenFailed string = "生成Token失败"

	AddScanStatusDescFailed string = "新建扫描失败"
	AddScanFailed           string = "执行扫描失败"
	ListScanFailed          string = "获取扫描列表失败"
	GetScanResultFailed     string = "获取扫描结果失败"

	ListAgentInfoFailed   string = "获取agent列表失败"
	DeleteAgentInfoFailed string = "删除agent失败"

	GetAgentInfoDetailsFailed string = "获取agent详情失败"

	PermissionDeny string = "权限不足"

	AddSecGrpRuleFailed    string = "新增安全组规则失败"
	DeleteSecGrpRuleFailed string = "删除安全组规则失败"
	ListSecGrpRuleFailed   string = "获取安全组规则列表失败"
	AddRepeatSecGrpRule    string = "新增规则重复"
	SecGrpRuleNotExist     string = "安全组规则不存在"

	PortRangeOverflow string = "端口无效"
	CIDRInvalid       string = "授权对象无效"

	StartBaselineScanFailed     string = "开启基线扫描失败"
	BaselineScanFailed          string = "基线扫描失败"
	GetBaselineScanResultFailed string = "获取基线扫描结果失败"
	UpdateBaselineStatusFailed  string = "更新基线状态失败"
	UpdateBaselineInvalid       string = "更新基线状态无效"
	GetBaselineInfoFailed       string = "获取基线规则详情失败"

	Unknow                           string = "未知..."
	ServerCreateChatFailed           string = "创建聊天服务失败"
	ServerReceiveChatFailed          string = "接收聊天服务失败"
	ClientNotSupportSSE              string = "客户端不支持SSE"
	ServerGetSessionMessageFailed    string = "获取会话消息失败"
	ServerAddSessionMessageFailed    string = "添加会话消息失败"
	ServerUpdateSessionMessageFailed string = "更新会话消息失败"
	ServerDeleteSessionMessageFailed string = "删除会话消息失败"
	GetAudioFileFailed               string = "获取音频文件失败"
	ServerTranscriptionFailed        string = "音频转换失败"
	ServerGenerateImageFailed        string = "生成图片失败"
	ServerGenerateAudioFailed        string = "生成音频失败"
	ServerGetAllSessionsFailed       string = "获取会话失败"
	ServerGetSessionMessagesFailed   string = "获取会话消息失败"
	ServerCreateSessionFailed        string = "新建会话失败"
	ServerUpdateSessionFailed        string = "更新会话失败"
	ChatSessionNotExist              string = "会话不存在"
)
