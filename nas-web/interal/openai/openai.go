package openai

import (
	"nas-web/config"
	"net/http"
	"net/url"

	"github.com/sashabaranov/go-openai"
)

var Client *openai.Client

func OpenaiInit(conf config.OpenaiConfig) {
	config := openai.DefaultConfig(conf.OpenaiApiKey)
	proxyUrl, err := url.Parse(conf.ProxyUrl)
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	config.HTTPClient = &http.Client{
		Transport: transport,
	}

	Client = openai.NewClientWithConfig(config)
}
