package config

type Config struct {
	OpenAIConfig   OpenAIConfig   `yaml:"openai"`
	PineconeConfig PineconeConfig `yaml:"pinecone"`
	DatabaseConfig DBConfig       `yaml:"database"`
}

type OpenAIConfig struct {
	APIKey        string  `yaml:"api_key"`
	MaxTokens     int     `yaml:"max_tokens"`
	Model         string  `yaml:"model"`
	ModelN        int     `yaml:"model_n"`
	Prompt        string  `yaml:"prompt"`
	Stream        bool    `yaml:"stream"`
	SystemContext string  `yaml:"system_context"`
	Temperature   float32 `yaml:"temperature"`
}

type PineconeConfig struct {
	APIKey      string `yaml:"api_key"`
	Environment string `yaml:"environment"`
	IndexName   string `yaml:"index_name"`
	ProjectName string `yaml:"project_name"`
}

type DBConfig struct {
	Name     string `yaml:"name"`
	Url      string `yaml:"url"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password`
}
