package service

import (
	"errors"
	"io"
	"nas-common/mlog"
	"nas-common/models"
	"nas-web/dao/ai"
	formjson "nas-web/dao/form_json"
	"nas-web/dao/mongo"
	"nas-web/interal/wrapper"
	"nas-web/support"
	webutils "nas-web/web-utils"
	"strconv"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
	"go.uber.org/zap"
)

// SendNoContextStreamChatHandler 发送无上下文流式聊天
func SendNoContextStreamChatHandler(ctx *wrapper.Context, reqBody interface{}) error {
	req := reqBody.(*formjson.SendNoContextStreamChatReq)

	stream, err := ai.Chat.RunWithNoContextStream(req.Question)
	if err != nil {
		mlog.Error("create no context stream chat failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ServerCreateChatFailed, 0)
		return nil
	}
	defer stream.Close()

	flusher, ok := ctx.ResponseWriter().Flusher()
	if !ok {
		mlog.Error("client not support SSE")
		support.SendApiErrorResponse(ctx, support.ClientNotSupportSSE, 0)
		return errors.New("client not support SSE")
	}
	ctx.ContentType("text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Transfer-Encoding", "chunked")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			ctx.Writef("event: the end of stream\n")
			return nil
		}

		if err != nil {
			mlog.Error("receive no context stream chat failed", zap.Error(err))
			support.SendApiErrorResponse(ctx, support.ServerReceiveChatFailed, 0)
			return nil
		}

		// must end with \n\n
		ctx.Writef("data: %s\n\n", response.Choices[0].Delta.Content)
		flusher.Flush()
	}
}

// SendContextStreamChatHandler 发送上下文流式聊天
func SendContextStreamChatHandler(ctx *wrapper.Context, reqBody interface{}) error {
	req := reqBody.(*formjson.SendContextStreamChatReq)
	resp := formjson.StatusResp{Status: "OK"}

	// [1]: check whether session exist
	usid := webutils.String.Hash(strconv.Itoa(ctx.UserToken.UserId), strconv.Itoa(req.SessionId))
	sessionMessagesDesc, err := mongo.Chat.GetByUSid(ctx, usid)
	if err != nil {
		mlog.Error("session not exist", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ChatSessionNotExist, 0)
		return nil
	}

	// [2]: create stream to run
	stream, err := ai.Chat.RunWithContextStream(req.Question, &sessionMessagesDesc)
	if err != nil {
		mlog.Error("create context stream chat failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ServerCreateChatFailed, 0)
		return nil
	}
	defer stream.Close()

	// [3]: check client support SSE
	flusher, ok := ctx.ResponseWriter().Flusher()
	if !ok {
		mlog.Error("client not support SSE")
		support.SendApiErrorResponse(ctx, support.ClientNotSupportSSE, 0)
		return errors.New("client not support SSE")
	}

	output := make([]string, 1000)

	// [4]: set SSE header
	ctx.ContentType("text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Transfer-Encoding", "chunked")

	// [5]: receive stream data
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			ctx.Writef("event: the end of stream\n")
			support.SendApiResponse(ctx, resp, "")
			break
		}

		if err != nil {
			mlog.Error("receive context stream chat failed", zap.Error(err))
			support.SendApiErrorResponse(ctx, support.ServerReceiveChatFailed, 0)
			return nil
		}

		ctx.Writef("data: %s\n\n", response.Choices[0].Delta.Content)
		output = append(output, response.Choices[0].Delta.Content)
		flusher.Flush()
	}

	// [6]: Update session message to db
	sessionMessages := make([]models.SessionMessages, 0, 2)
	sessionMessages = append(sessionMessages, models.SessionMessages{
		Role:    support.ChatMessageRoleUser,
		Content: req.Question,
	}, models.SessionMessages{
		Role:    support.ChatMessageRoleAssistant,
		Content: strings.Join(output, ""),
	},
	)
	change := bson.M{"$push": bson.M{"messages": bson.M{"$each": sessionMessages}}} // append to inner message array
	if err = mongo.Chat.AppendMessages(ctx, usid, change); err != nil {
		mlog.Error("update session message failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ServerUpdateSessionMessageFailed, 0)
		return nil
	}
	return nil
}

// DeleteContextStreamChatHandler 删除上下文流式聊天
func DeleteContextStreamChatHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.DeleteContextStreamChatReq)
	resp := formjson.StatusResp{Status: "OK"}

	usid := webutils.String.Hash(strconv.Itoa(ctx.UserToken.UserId), strconv.Itoa(req.SessionId))
	err = mongo.Chat.DeleteSession(ctx, usid)
	if err != nil {
		mlog.Error("delete session failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ServerDeleteSessionMessageFailed, 0)
		return nil
	}
	support.SendApiResponse(ctx, resp, "")

	return nil
}

// GetAllSessionsHandler 获取指定用户的所有会话
func GetAllSessionsHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	var sessionMessagesDescs []models.SessionMessagesDesc
	if sessionMessagesDescs, err = mongo.Chat.GetAllSessionsByUid(ctx, ctx.UserToken.UserId); err != nil {
		mlog.Error("get all sessions failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ServerGetAllSessionsFailed, 0)
		return nil
	}
	datas := make([]formjson.SessionData, 0, len(sessionMessagesDescs))
	for _, sessionMessagesDesc := range sessionMessagesDescs {
		datas = append(datas, formjson.SessionData{
			SessionId:   sessionMessagesDesc.SessionId,
			SessionName: sessionMessagesDesc.SessionName,
			CreateTime:  sessionMessagesDesc.StartTime.Unix(),
		})
	}

	resp := formjson.GetAllSessionsResp{
		Uid:   ctx.UserToken.UserId,
		Datas: datas,
	}
	support.SendApiResponse(ctx, resp, "")
	return
}

// GetSessionMessagesHandler 获取指定会话的所有消息
func GetSessionMessagesHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.GetSessionMessagesReq)
	usid := webutils.String.Hash(strconv.Itoa(ctx.UserToken.UserId), strconv.Itoa(req.SessionId))
	var doc models.SessionMessagesDesc
	if doc, err = mongo.Chat.GetByUSid(ctx, usid); err != nil {
		mlog.Error("get session messages failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ServerGetSessionMessageFailed, 0)
		return nil
	}
	messages := make([]formjson.SessionMessages, 0, len(doc.Messages))
	for _, message := range doc.Messages {
		messages = append(messages, formjson.SessionMessages{
			Role:    message.Role,
			Content: message.Content,
		})
	}

	resp := formjson.GetSessionMessagesResp{
		Uid:         ctx.UserToken.UserId,
		Model:       doc.Model,
		Messages:    messages,
		MaxTokens:   doc.MaxTokens,
		Temperature: doc.Temperature,
		SessionName: doc.SessionName,
		System:      doc.System,
	}
	support.SendApiResponse(ctx, resp, "")
	return
}

// CreateSessionHandler 新建会话
func CreateSessionHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.CreateSessionReq)
	resp := formjson.CreateSessionResp{}
	newSessionId := mongo.Chat.GetMaxSessionId(ctx, ctx.UserToken.UserId)

	sessionMessagesDesc := models.SessionMessagesDesc{
		USid:        webutils.String.Hash(strconv.Itoa(ctx.UserToken.UserId), strconv.Itoa(newSessionId)),
		Uid:         ctx.UserToken.UserId,
		SessionId:   newSessionId,
		StartTime:   time.Now(),
		Model:       req.Model,
		Temperature: req.Temperature,
		MaxTokens:   req.MaxTokens,
		Stop:        req.Stop,
		SessionName: req.SessionName,
		System:      req.System,
	}

	if err = mongo.Chat.AddSession(ctx, &sessionMessagesDesc); err != nil {
		mlog.Error("create session failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ServerCreateSessionFailed, 0)
		return nil
	}

	resp.SessionId = newSessionId
	support.SendApiResponse(ctx, resp, "")
	return
}

// UpdateSessionHandler 更新会话
func UpdateSessionHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.UpdateSessionReq)
	resp := formjson.UpdateSessionResp{SessionId: req.SessionId}

	usid := webutils.String.Hash(strconv.Itoa(ctx.UserToken.UserId), strconv.Itoa(req.SessionId))

	update := bson.M{
		"max_tokens":   req.MaxTokens,
		"temperature":  req.Temperature,
		"stop":         req.Stop,
		"model":        req.Model,
		"session_name": req.SessionName,
		"system":       req.System,
	}

	if err = mongo.Chat.UpdateModelParams(ctx, usid, update); err != nil {
		mlog.Error("update session failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ServerUpdateSessionFailed, 0)
		return nil
	}

	support.SendApiResponse(ctx, resp, "")
	return
}

// DeleteSessionMessagesHandler 删除指定会话的所有消息
func DeleteSessionMessagesHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.DeleteSessionMessagesReq)
	resp := formjson.StatusResp{Status: "OK"}
	usid := webutils.String.Hash(strconv.Itoa(ctx.UserToken.UserId), strconv.Itoa(req.SessionId))
	empty := make([]models.SessionMessages, 0, 0)
	change := bson.M{"$set": bson.M{"messages": empty}} // delete to only keep system messgae
	if err = mongo.Chat.DeleteMessages(ctx, usid, change); err != nil {
		mlog.Error("delete session messages failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ServerDeleteSessionMessageFailed, 0)
		return nil
	}
	support.SendApiResponse(ctx, resp, "")
	return
}

// DeleteAllSessionsHandler 删除指定用户的所有会话
func DeleteAllSessionsHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	resp := formjson.StatusResp{Status: "OK"}
	if err = mongo.Chat.DeleteAllSessions(ctx, ctx.UserToken.UserId); err != nil {
		mlog.Error("delete all sessions failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ServerDeleteAllSessionFailed, 0)
		return nil
	}
	support.SendApiResponse(ctx, resp, "")
	return
}
