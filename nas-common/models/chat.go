package models

import "time"

type SessionMessages struct {
	Role    string `bson:"role"`    //消息角色
	Content string `bson:"content"` //消息内容
}

type SessionMessagesDesc struct {
	USid        string            `bson:"_id"`          // hash(用户id + ":" + 会话id)
	Uid         int               `bson:"uid"`          // 用户id
	SessionId   int               `bson:"session_id"`   // 会话id
	SessionName string            `bson:"session_name"` // 会话名称
	StartTime   time.Time         `bson:"start_time"`   // 会话创建时间
	Model       string            `bson:"model"`        // 模型名称
	MaxTokens   int               `bson:"max_tokens"`   // 输入+输出的最大长度
	Temperature float32           `bson:"temperature"`  // 生成文本的多样性
	System      string            `bson:"system"`       // system message
	Stop        []string          `bson:"stop"`         // 匹配到这些词时停止生成
	Messages    []SessionMessages `bson:"messages"`     // messages
}

func (SessionMessagesDesc) Collection() string {
	return "tb_chat_session"
}
