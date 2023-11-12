package service

import (
	"nas-common/mlog"
	formjson "nas-web/dao/form_json"
	"nas-web/interal/wrapper"
	"nas-web/support"
	webutils "nas-web/web-utils"
	"path/filepath"

	"go.uber.org/zap"
)

// TranscriptionsHandler transcribes audio into the input language
func TranscriptionsHandler(ctx *wrapper.Context, reqBody interface{}) error {
	req := reqBody.(*formjson.TranscriptionsReq)
	resp := formjson.TranscriptionsResp{}
	file, fileHeader, err := ctx.FormFile("file")
	if err != nil {
		mlog.Error("get audio file failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.GetAudioFileFailed, 0)
		return err
	}
	defer file.Close()
	fullPath := filepath.Join("/home/aico/Music", fileHeader.Filename)
	err = webutils.System.SaveFile(file, fullPath)

	// resp.Answer, err = ai.Chat.RunWithNoContextNoStream(req.ModelName, req.Question)
	// if err != nil {
	// 	mlog.Error("create no context no stream chat failed", zap.Error(err))
	// 	support.SendApiErrorResponse(ctx, support.ServerCreateChatFailed, 0)
	// }

	resp.Answer = "upload ok, req lang is " + req.Language
	support.SendApiResponse(ctx, resp, "")
	return nil
}
