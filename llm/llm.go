package llm

import (
	"bufio"
	"context"
	"fmt"
	"github.com/aldarisbm/ltmllm/config"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"os"
	"strings"
)

type ChatBot struct {
	client *openai.Client
	Cfg    *config.Config
}

func NewChatBot(cfg *config.Config) ChatBot {
	c := openai.NewClient(cfg.OpenAIConfig.APIKey)
	return ChatBot{
		client: c,
		Cfg:    cfg,
	}
}

func (b *ChatBot) Chat() {
	for {
		input := b.getStdInput()
		if input == "quit" {
			break
		}
		b.sendRequest(input)
	}
}

func (b *ChatBot) sendRequest(input string) {
	ctx := context.Background()
	// TODO maybe add user here to track who said what
	messages := []openai.ChatCompletionMessage{
		{
			Content: b.Cfg.OpenAIConfig.SystemContext,
			Role:    openai.ChatMessageRoleSystem,
		},
		{
			Content: input,
			Role:    openai.ChatMessageRoleUser,
		},
	}

	req := openai.ChatCompletionRequest{
		Model:       b.Cfg.OpenAIConfig.Model,
		Messages:    messages,
		Temperature: b.Cfg.OpenAIConfig.Temperature,
		TopP:        1,
		N:           b.Cfg.OpenAIConfig.ModelN,
		Stream:      b.Cfg.OpenAIConfig.Stream,
	}
	if err := b.processRequest(ctx, req); err != nil {
		log.Panic(err)
	}
}

func (b *ChatBot) processRequest(ctx context.Context, req openai.ChatCompletionRequest) error {
	stream, err := b.client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
	}
	defer stream.Close()
	b.processStream(stream)
	return nil
}

func (b *ChatBot) processStream(stream *openai.ChatCompletionStream) {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Print(resp.Choices[0].Delta.Content)
	}
}

func (b *ChatBot) getStdInput() string {
	fmt.Print(b.Cfg.OpenAIConfig.Prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	trimmedInput := strings.TrimSpace(input)

	return trimmedInput
}
