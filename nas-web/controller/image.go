package controller

import (
	formjson "nas-web/dao/form_json"
	"nas-web/interal/wrapper"
	"nas-web/service"
	"nas-web/support"
)

type ImageController struct{}

// @Summary generate image
// @Description generate image by using text
// @Tags Image
// @Accept json
// @Produce json
// @Param auth body formjson.GenerateImageReq true "request data"
// @Success 200 {object} formjson.GenerateImageResp "response data"
// @Router /v1/image/generate [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (a ImageController) GenerateImage(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.GenerateHandler, true, &formjson.GenerateImageReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}
