package ai

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	formjson "nas-web/dao/form_json"
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

func (a *audio) Speech(req *formjson.SpeechReq) (io.ReadCloser, error) {
	// check model name
	modelName := openai.SpeechModel(req.ModelName)
	if modelName != openai.TTSModel1 && modelName != openai.TTsModel1HD {
		modelName = openai.TTSModel1
	}
	// check Input
	if req.Input == "" {
		return nil, errors.New("Input text is empty.")
	}
	// check voice
	// TODO

	resp, err := innerOpenai.Client.CreateSpeech(context.Background(), openai.CreateSpeechRequest{
		Model:          modelName,
		Input:          req.Input,
		Voice:          openai.SpeechVoice(req.Voice),
		Speed:          req.Speed,
		ResponseFormat: openai.SpeechResponseFormat(req.Format),
	})
	return resp, err
}
