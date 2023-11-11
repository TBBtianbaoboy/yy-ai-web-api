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

func (c *chat) RunWithNoContextNoStream(modelName string, question string) (string, error) {
	resp, err := innerOpenai.Client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: modelName,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: question,
			},
		},
	})
	return resp.Choices[0].Message.Content, err
}

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

func (c *chat) RunWithContextStream(modelName string, question string, sessionMessagesDesc *models.SessionMessagesDesc) (*openai.ChatCompletionStream, error) {
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
		Model:    modelName,
		Messages: messages,
		Stream:   true,
	})
	return stream, err
}
