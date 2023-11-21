package ai

import (
	"context"
	"fmt"
	"nas-common/models"
	innerOpenai "nas-web/interal/openai"

	"github.com/sashabaranov/go-openai"
)

type chat struct{}

var Chat chat

func (c *chat) RunWithNoContextStream(modelName string, question string) (*openai.ChatCompletionStream, error) {
	stream, err := innerOpenai.Client.CreateChatCompletionStream(context.Background(), openai.ChatCompletionRequest{
		Model: modelName,
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
	messages := make([]openai.ChatCompletionMessage, len(sessionMessagesDesc.Messages), len(sessionMessagesDesc.Messages)+1)
	for i, v := range sessionMessagesDesc.Messages {
		messages[i].Role = v.Role
		messages[i].Content = v.Content
	}
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: question,
	})
	for i, v := range messages {
		fmt.Println(i, v.Role, v.Content)
	}
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
