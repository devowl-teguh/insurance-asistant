package llm

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

type NaraRouter struct {
	client openai.Client
	model  string
}

func NewNaraRouter(apiKey, model string) *NaraRouter {
	return newNaraRouter(apiKey, model, "https://router.bynara.id/v1/")
}

func newNaraRouter(apiKey, model, baseURL string) *NaraRouter {
	return &NaraRouter{client: openai.NewClient(option.WithBaseURL(baseURL), option.WithAPIKey(apiKey), option.WithMaxRetries(0), option.WithRequestTimeout(45*time.Second)), model: model}
}
func (n *NaraRouter) Answer(ctx context.Context, question string) (string, error) {
	response, err := n.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: n.model,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage("You are a helpful insurance learning assistant. Answer general questions clearly. You do not have access to any policy or customer information. Do not claim otherwise. If asked about a specific policy, explain that you need its text."),
			openai.UserMessage(question),
		},
	})
	if err != nil {
		return "", err
	}
	if len(response.Choices) == 0 || strings.TrimSpace(response.Choices[0].Message.Content) == "" {
		return "", errors.New("model returned an empty answer")
	}
	return response.Choices[0].Message.Content, nil
}
