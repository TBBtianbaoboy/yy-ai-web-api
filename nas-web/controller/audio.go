package controller

import (
	formjson "nas-web/dao/form_json"
	"nas-web/interal/wrapper"
	"nas-web/service"
	"nas-web/support"
)

type AudioController struct{}

// @Summary Transcribes audio into the input language
// @Description Transcribes audio into the input language
// @Tags Audio
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "audio file data"
// @Param auth query formjson.TranscriptionsReq true "request data"
// @Success 200 {object} formjson.TranscriptionsResp "response data"
// @Router /v1/audio/transcriptions [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a AudioController) Transcriptions(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.TranscriptionsHandler, true, &formjson.TranscriptionsReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

// @Summary Generate audio from the input text
// @Description Generate audio from the input text
// @Tags Audio
// @Accept json
// @Produce octet-stream
// @Param auth body formjson.SpeechReq true "request data"
// @Success 200 {file} - "response data"
// @Router /v1/audio/speech [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a AudioController) Speech(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.SpeechHandler, true, &formjson.SpeechReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}
