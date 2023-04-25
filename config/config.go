package config

import "os"
import "gopkg.in/yaml.v2"

type Config struct {
	OpenAIKey   string `yaml:"open_ai_key"`
	OpenAIModel string `yaml:"open_ai_model"`

	DBName     string `yaml:"db_name"`
	DBUrl      string `yaml:"db_url"`
	DBPort     string `yaml:"db_port"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
}

func NewConfig() *Config {
	f, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	// loading secrets from env vars
	cfg.DBPassword = os.Getenv("OPENAI_KEY")
	cfg.OpenAIKey = os.Getenv("DB_PASSWORD")

	return &cfg
}
