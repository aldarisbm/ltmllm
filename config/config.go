package config

import "os"
import "gopkg.in/yaml.v2"

func NewConfig() *Config {
	f, err := os.Open(PropertiesFile)
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
	cfg.OpenAIConfig.APIKey = os.Getenv("OPENAI_KEY")
	cfg.DatabaseConfig.Password = os.Getenv("DB_PASSWORD")
	cfg.PineconeConfig.APIKey = os.Getenv("PINECONE_KEY")
	return &cfg
}
