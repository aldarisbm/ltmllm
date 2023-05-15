package chatbot

import (
	"bufio"
	"context"
	"fmt"
	memory "github.com/aldarisbm/memory/pkg"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"os"
	"strings"
)

type ChatBot struct {
	apiKey        string
	systemContext string
	model         string
	temperature   float32
	n             int
	prompt        string
	topP          float32
	stream        bool

	client *openai.Client
	memory *memory.Memory
}

func (b *ChatBot) StdInput() {
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

func (b *ChatBot) ChatStream(input string) *openai.ChatCompletionStream {
	req := b.getRequest(input)
	stream, err := b.processRequest(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	return stream
}

func (b *ChatBot) getRequest(input string) openai.ChatCompletionRequest {
	// TODO maybe add user here to track who said what
	doc := b.memory.NewDocument(input, "aldarisbm")
	if err := b.memory.StoreDocument(doc); err != nil {
		// should probably not be a panic
		panic(err)
	}

	messages := []openai.ChatCompletionMessage{
		{
			Content: b.systemContext,
			Role:    openai.ChatMessageRoleSystem,
		},
		{
			Content: input,
			Role:    openai.ChatMessageRoleUser,
		},
	}

	req := openai.ChatCompletionRequest{
		Model:       b.model,
		Messages:    messages,
		Temperature: b.temperature,
		TopP:        b.topP,
		N:           b.n,
		Stream:      b.stream,
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
	var sb strings.Builder
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			document := b.memory.NewDocument(sb.String(), "chatbot")
			if err := b.memory.StoreDocument(document); err != nil {
				panic(err)
			}
			break
		}
		sb.WriteString(resp.Choices[0].Delta.Content)
		fmt.Print(resp.Choices[0].Delta.Content)
	}
}

func (b *ChatBot) getStdInput() string {
	fmt.Print(b.prompt)
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
