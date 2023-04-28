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

func (b *ChatBot) ChatStdInput() {
	for {
		input := b.getStdInput()
		if input == "quit" {
			break
		}
		req := b.getRequest(input)
		stream, err := b.processRequest(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		b.processStream(stream)
		fmt.Println()
	}
}

func (b *ChatBot) Chat(input string) string {
	req := b.getRequest(input)
	stream, err := b.processRequest(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	return b.processStreamToString(stream)
}

func (b *ChatBot) getRequest(input string) openai.ChatCompletionRequest {
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
	return req
}

func (b *ChatBot) processRequest(ctx context.Context, req openai.ChatCompletionRequest) (*openai.ChatCompletionStream, error) {
	stream, err := b.client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("ChatCompletionStream error: %v\n", err)
	}
	return stream, err
}

func (b *ChatBot) processStream(stream *openai.ChatCompletionStream) {
	defer stream.Close()
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

func (b *ChatBot) processStreamToString(stream *openai.ChatCompletionStream) string {
	defer stream.Close()
	var sb strings.Builder
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		sb.WriteString(resp.Choices[0].Delta.Content)
		fmt.Print(resp.Choices[0].Delta.Content)
	}
	return sb.String()
}
