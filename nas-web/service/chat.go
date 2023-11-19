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

// SendNoContextNoStreamChatHandler 发送无上下文无流式聊天
func SendNoContextNoStreamChatHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.SendNoContextNoStreamChatReq)
	resp := formjson.SendNoContextNoStreamChatResp{}

	resp.Answer, err = ai.Chat.RunWithNoContextNoStream(req.ModelName, req.Question)
	if err != nil {
		mlog.Error("create no context no stream chat failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ServerCreateChatFailed, 0)
		return
	}

	support.SendApiResponse(ctx, resp, "")
	return
}

// SendNoContextStreamChatHandler 发送无上下文流式聊天
func SendNoContextStreamChatHandler(ctx *wrapper.Context, reqBody interface{}) error {
	req := reqBody.(*formjson.SendNoContextStreamChatReq)

	stream, err := ai.Chat.RunWithNoContextStream(req.ModelName, req.Question)
	if err != nil {
		mlog.Error("create no context stream chat failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ServerCreateChatFailed, 0)
		return err
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
			return err
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
	sessionMessagesDesc, err, exist := mongo.Chat.GetByUSid(ctx, usid)
	if err != nil {
		mlog.Error("get session message failed", zap.Error(err))
		// support.SendApiErrorResponse(ctx, support.ServerGetSessionMessageFailed, 0)
		// return err
	}

	// [2]: create stream to run
	stream, err := ai.Chat.RunWithContextStream(req.ModelName, req.Question, &sessionMessagesDesc)
	if err != nil {
		mlog.Error("create context stream chat failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ServerCreateChatFailed, 0)
		return err
	}
	defer stream.Close()

	// [3]: check client support SSE
	flusher, ok := ctx.ResponseWriter().Flusher()
	if !ok {
		mlog.Error("client not support SSE")
		support.SendApiErrorResponse(ctx, support.ClientNotSupportSSE, 0)
		return errors.New("client not support SSE")
	}

	output := make([]string, 100)

	// [4]: set SSE header
	ctx.ContentType("text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	// [5]: receive stream data
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			ctx.Writef("event: the end of stream\n")
			support.SendApiResponse(ctx, resp, "")
			break
		}

		if err != nil {
			mlog.Error("receive no context stream chat failed", zap.Error(err))
			support.SendApiErrorResponse(ctx, support.ServerReceiveChatFailed, 0)
			return err
		}

		ctx.Writef("data: %s\n", response.Choices[0].Delta.Content)
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
	if exist {
		change := bson.M{"$push": bson.M{"messages": bson.M{"$each": sessionMessages}}} // append to inner message array
		if err = mongo.Chat.AppendMessages(ctx, usid, change); err != nil {
			mlog.Error("update session message failed", zap.Error(err))
			support.SendApiErrorResponse(ctx, support.ServerUpdateSessionMessageFailed, 0)
			return err
		}
	} else {
		sessionMessagesDesc = models.SessionMessagesDesc{
			USid:      usid,
			Uid:       ctx.UserToken.UserId,
			SessionId: req.SessionId,
			StartTime: time.Now(),
			Messages:  sessionMessages,
		}
		if err = mongo.Chat.AddSession(ctx, &sessionMessagesDesc); err != nil {
			mlog.Error("add session message failed", zap.Error(err))
			support.SendApiErrorResponse(ctx, support.ServerAddSessionMessageFailed, 0)
			return err
		}
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
		return
	}
	support.SendApiResponse(ctx, resp, "")

	return
}

// GetAllSessionsHandler 获取指定用户的所有会话
func GetAllSessionsHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	var userDoc models.User
	if err = mongo.User.FindByUid(ctx, ctx.UserToken.UserId, &userDoc); err != nil {
		mlog.Error("get current user info failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.GetUserInfoFailed, 0)
		return nil
	}
	resp := formjson.GetUserInfoResp{
		Uid:        userDoc.UID,
		UserType:   userDoc.UserType,
		UserName:   userDoc.UserName,
		Email:      userDoc.Mail,
		Phone:      userDoc.Mobile,
		PS:         userDoc.PasswordStrength,
		CreateTime: userDoc.InsertTm.Unix(),
	}
	support.SendApiResponse(ctx, resp, "")
	return
}
