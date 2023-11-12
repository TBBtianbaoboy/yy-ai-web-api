package models

import "time"

type SessionMessages struct {
	Role    string `bson:"role"`    //消息角色
	Content string `bson:"content"` //消息内容
}
type SessionMessagesDesc struct {
	USid      string            `bson:"_id"`        // session id MD5(user id + session id)
	Uid       int               `bson:"uid"`        // user id
	SessionId int               `bson:"session_id"` // session id
	StartTime time.Time         `bson:"start_time"` // start time
	Messages  []SessionMessages `bson:"messages"`   // messages
}

func (SessionMessagesDesc) Collection() string {
	return "tb_chat_session"
}
