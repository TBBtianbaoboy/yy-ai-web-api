package service

import (
	"nas-common/mlog"
	"nas-web/dao/ai"
	formjson "nas-web/dao/form_json"
	"nas-web/interal/wrapper"
	"nas-web/support"

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

	resp.Text, err = ai.Audio.Transcriptions(file, fileHeader.Filename, req.Language)
	if err != nil {
		mlog.Error("audio transcriptions failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ServerTranscriptionFailed, 0)
		return err
	}

	support.SendApiResponse(ctx, resp, "")
	return nil
}
