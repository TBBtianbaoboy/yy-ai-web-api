package service

import (
	"io"
	"nas-common/mlog"
	"nas-web/dao/ai"
	formjson "nas-web/dao/form_json"
	"nas-web/interal/wrapper"
	"nas-web/support"
	webutils "nas-web/web-utils"

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

// SpeechHandler generate audio from the input text
func SpeechHandler(ctx *wrapper.Context, reqBody interface{}) error {
	req := reqBody.(*formjson.SpeechReq)

	audio_file, err := ai.Audio.Speech(req)
	if err != nil {
		mlog.Error("generate audio from input text failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ServerGenerateAudioFailed, 0)
		return err
	}
	defer audio_file.Close()

	ctx.ContentType("application/octet-stream")

	fileName := webutils.String.GetRandomString(10) + ".mp3"
	ctx.Header("Content-Disposition", "attachment;filename="+fileName)

	ctx.StreamWriter(func(w io.Writer) bool {
		size, err := io.Copy(w, audio_file)
		if err != nil {
			mlog.Error("copy audio file failed", zap.Error(err))
			return false
		}
		if size == 0 {
			return false
		}
		return true
	})

	return nil
}
