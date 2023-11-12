package service

import (
	"nas-common/mlog"
	"nas-web/dao/ai"
	formjson "nas-web/dao/form_json"
	"nas-web/interal/wrapper"
	"nas-web/support"

	"go.uber.org/zap"
)

// GenerateHandler generate image by using prompt
func GenerateHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	req := reqBody.(*formjson.GenerateImageReq)
	resp := formjson.GenerateImageResp{}

	resp.Base64, err = ai.Image.Generate(req)
	if err != nil {
		mlog.Error("generate image failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.ServerGenerateImageFailed, 0)
		return
	}

	support.SendApiResponse(ctx, resp, "")
	return
}
