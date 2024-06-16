package config

import (
	"os"
)
import "gopkg.in/yaml.v2"

// NewConfig loads the configuration from the properties file
// and returns a config struct.
func NewConfig() config {
	propertiesFile := os.Getenv("PROPERTIES_FILE")
	if propertiesFile == "" {
		propertiesFile = DefaultPropertiesFile
	}

	f, err := os.Open(propertiesFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	// loading secrets from env vars
	cfg.OpenAIConfig.APIKey = os.Getenv("OPENAI_API_KEY")
	cfg.DatabaseConfig.Password = os.Getenv("DB_PASSWORD")
	cfg.PineconeConfig.APIKey = os.Getenv("PINECONE_KEY")
	return cfg
}
