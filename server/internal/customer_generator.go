package internal

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type CustomerGenerator struct {
	Prompt string

	BaseURL  string
	APIKey   string
	Model    string
	ProxyURL string
}

func (c CustomerGenerator) Generate(query string) (string, error) {

	// create client
	conf := openai.DefaultConfig(c.APIKey)
	if c.BaseURL != "" {
		conf.BaseURL = c.BaseURL
	}
	if c.ProxyURL != "" {
		proxyUrl, err := url.Parse(c.ProxyURL)
		if err != nil {
			return "", fmt.Errorf("create OpenaiChatGPT client failed on parsing proxy url, err: %v", err)
		}
		conf.HTTPClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	}
	client := openai.NewClientWithConfig(conf)

	messages := append(
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
	)

	slog.Info("chat", slog.Any("messages", messages))

	// create request
	req := openai.ChatCompletionRequest{
		Messages: messages,
		Model:    c.Model,
		Stream:   false,
	}
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("CreateChatCompletion failed, err %v", err)
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("empty resp %v", resp)
	}

	content := resp.Choices[0].Message.Content

	// remove ```json and ```, only keep json contents
	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimSuffix(content, "```")
	slog.Info("customer generated", slog.String("content", content))
	return content, nil
}
