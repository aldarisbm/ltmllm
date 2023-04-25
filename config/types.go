package config

type Config struct {
	OpenAIConfig   OpenAIConfig `yaml:"openai"`
	DatabaseConfig DBConfig     `yaml:"database"`
}

type OpenAIConfig struct {
	Model       string  `yaml:"model"`
	APIKey      string  `yaml:"api_key,omitempty"`
	ModenN      int     `yaml:"model_n"`
	MaxTokens   int     `yaml:"max_tokens"`
	Temperature float64 `yaml:"temperature"`
	Stream      bool    `yaml:"stream"`
}

type DBConfig struct {
	Name     string `yaml:"name"`
	Url      string `yaml:"url"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
