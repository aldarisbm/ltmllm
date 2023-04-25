package llm

//
//import (
//	"bufio"
//	"fmt"
//	"github.com/sashabaranov/go-openai"
//	"log"
//	"os"
//	"strings"
//)
//
//type ChatBot struct {
//	client        *openai.Client
//	model         string
//	systemContext string
//	config        *Config
//}
//
//type Config struct {
//	N           int
//	MaxTokens   int
//	Temperature float32
//	Stream      bool
//}
//
//func getInput(s string) string {
//	fmt.Print(s)
//	reader := bufio.NewReader(os.Stdin)
//	token, err := reader.ReadString('\n')
//	if err != nil {
//		log.Panic(err)
//	}
//	trimmedInput := strings.TrimSpace(token)
//
//	return trimmedInput
//}
//
//func msg() {
//	req := openai.ChatCompletionRequest{
//		Model: b.model,
//		Messages: []openai.ChatCompletionMessage{
//			{
//				Role:    "system",
//				Content: b.systemContext,
//			},
//		},
//		N:           1,
//		MaxTokens:   0,
//		Temperature: 0.5,
//		Stream:      true,
//	}
//}
