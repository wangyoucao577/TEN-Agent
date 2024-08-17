package internal

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/sashabaranov/go-openai"
)

type CustomerGenerator struct {
	Prompt string

	BaseURL  string
	APIKey   string
	Model    string
	ProxyURL string
}

func (c CustomerGenerator) Generate(query string) error {
	slog.Info("customer generate", slog.String("query", query))

	// create client
	conf := openai.DefaultConfig(c.APIKey)
	if c.BaseURL != "" {
		conf.BaseURL = c.BaseURL
	}
	if c.ProxyURL != "" {
		proxyUrl, err := url.Parse(c.ProxyURL)
		if err != nil {
			return fmt.Errorf("create OpenaiChatGPT client failed on parsing proxy url, err: %v", err)
		}
		conf.HTTPClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	}
	client := openai.NewClientWithConfig(conf)

	// create request
	req := openai.ChatCompletionRequest{
		Messages: append(
			[]openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: c.Prompt,
				},
			},
			openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: query,
			},
		),
		Model:  c.Model,
		Stream: false,
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return fmt.Errorf("CreateChatCompletion failed, err %v", err)
	}
	if len(resp.Choices) == 0 {
		return fmt.Errorf("empty resp %v", resp)
	}

	// TODO: return
	slog.Info("customer generated", slog.String("content", resp.Choices[0].Message.Content))
	return nil
}
