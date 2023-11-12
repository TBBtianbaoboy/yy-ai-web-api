package ai

import (
	"context"
	"mime/multipart"
	innerOpenai "nas-web/interal/openai"

	"github.com/sashabaranov/go-openai"
)

type audio struct{}

var Audio audio

func (a *audio) Transcriptions(file multipart.File, fullPath string, language string) (string, error) {
	resp, err := innerOpenai.Client.CreateTranscription(context.Background(), openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: fullPath,
		Reader:   file,
		Language: language,
	})
	return resp.Text, err
}
