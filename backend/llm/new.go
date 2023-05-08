package llm

import "github.com/sashabaranov/go-openai"

func NewChatBot(opts ...func(*ChatBot)) *ChatBot {
	const (
		systemContext = "You're a helpful assistant"
		model         = "gpt-3.5-turbo"
		temperature   = 0
		topP          = 1
		n             = 1
		stream        = true
		prompt        = "How can I help you today?: "
	)

	// topP, n,  and stream are set to default values always
	cb := ChatBot{
		systemContext: systemContext,
		model:         model,
		temperature:   temperature,
		n:             n,
		prompt:        prompt,
		topP:          topP,
		stream:        stream,
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

func WithPrompt(prompt string) func(*ChatBot) {
	return func(cb *ChatBot) {
		cb.prompt = prompt
	}
}
