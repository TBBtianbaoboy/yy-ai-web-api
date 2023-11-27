package ai

import (
	"context"
	"nas-common/models"
	innerOpenai "nas-web/interal/openai"

	"github.com/sashabaranov/go-openai"
)

type chat struct{}

var Chat chat

func (c *chat) RunWithNoContextStream(question string) (*openai.ChatCompletionStream, error) {
	stream, err := innerOpenai.Client.CreateChatCompletionStream(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo1106,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: question,
			},
		},
		Stream: true,
	})
	return stream, err
}

func (c *chat) RunWithContextStream(question string, sessionMessagesDesc *models.SessionMessagesDesc) (*openai.ChatCompletionStream, error) {
	messages := make([]openai.ChatCompletionMessage, len(sessionMessagesDesc.Messages)+1, len(sessionMessagesDesc.Messages)+2)
	messages[0].Role = openai.ChatMessageRoleSystem
	messages[0].Content = sessionMessagesDesc.System
	for i, v := range sessionMessagesDesc.Messages {
		messages[i+1].Role = v.Role
		messages[i+1].Content = v.Content
	}
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: question,
	})
	stream, err := innerOpenai.Client.CreateChatCompletionStream(context.Background(), openai.ChatCompletionRequest{
		Model:       sessionMessagesDesc.Model,
		Messages:    messages,
		Stream:      true,
		Temperature: sessionMessagesDesc.Temperature,
		MaxTokens:   sessionMessagesDesc.MaxTokens,
		Stop: func() []string {
			tempStop := make([]string, 0, len(sessionMessagesDesc.Stop))
			for _, v := range sessionMessagesDesc.Stop {
				if v == "" {
					continue
				}
				tempStop = append(tempStop, v)
			}
			return tempStop
		}(),
	})
	return stream, err
}
