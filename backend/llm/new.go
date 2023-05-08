package llm

import "github.com/sashabaranov/go-openai"

func NewChatBot(opts ...func(*ChatBot)) *ChatBot {
	const (
		systemContext = "You're a helpful assistant"
		model         = "gpt-3.5-turbo"
		temperature   = 0.7
		topP          = 1.0
		n             = 1
		stream        = true
		prompt        = "Enter your message: "
	)

	cb := ChatBot{
		systemContext: systemContext,
		model:         model,
		temperature:   temperature,
		topP:          topP,
		n:             n,
		stream:        stream,
		prompt:        prompt,
	}

	for _, o := range opts {
		o(&cb)
	}
	if cb.apiKey == "" {
		panic("API key is required")
	}
	cb.client = openai.NewClient(cb.apiKey)
	return &cb
}

func WithApiKey(apiKey string) func(*ChatBot) {
	return func(cb *ChatBot) {
		cb.apiKey = apiKey
	}
}

func WithSystemContext(systemContext string) func(*ChatBot) {
	return func(cb *ChatBot) {
		cb.systemContext = systemContext
	}
}

func WithModel(model string) func(*ChatBot) {
	return func(cb *ChatBot) {
		cb.model = model
	}
}

func WithTemperature(temperature float32) func(*ChatBot) {
	return func(cb *ChatBot) {
		cb.temperature = temperature
	}
}

func WithTopP(topP float64) func(*ChatBot) {
	return func(cb *ChatBot) {
		cb.topP = topP
	}
}

func WithN(n int) func(*ChatBot) {
	return func(cb *ChatBot) {
		cb.n = n
	}
}

func WithStream(stream bool) func(*ChatBot) {
	return func(cb *ChatBot) {
		cb.stream = stream
	}
}

func WithPrompt(prompt string) func(*ChatBot) {
	return func(cb *ChatBot) {
		cb.prompt = prompt
	}
}
