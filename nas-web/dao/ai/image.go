package ai

import (
	"context"
	formjson "nas-web/dao/form_json"
	innerOpenai "nas-web/interal/openai"

	"github.com/sashabaranov/go-openai"
)

type image struct{}

var Image image

func (a *image) Generate(req *formjson.GenerateImageReq) (string, error) {
	resp, err := innerOpenai.Client.CreateImage(context.Background(), openai.ImageRequest{
		Model:          req.ModelName,
		Prompt:         req.Prompt,
		Size:           req.Size,
		N:              1,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		Quality:        req.Quality,
	})
	if err != nil {
		return "", err
	}
	return resp.Data[0].B64JSON, err
}
