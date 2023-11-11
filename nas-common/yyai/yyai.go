package yyai

import (
	"net/http"
	"net/url"

	"github.com/sashabaranov/go-openai"
)

func NewClientWithProxy(openApiKey string, rawUrl string) (*openai.Client, error) {
	config := openai.DefaultConfig(openApiKey)
	proxyUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	config.HTTPClient = &http.Client{
		Transport: transport,
	}

	return openai.NewClientWithConfig(config), nil
}
