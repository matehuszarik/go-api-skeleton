package config

import (
	"log"

	// ...
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

// Configuration ..
type Configuration struct {
	Port    int      `envconfig:"port" default:"80"`
	APIKeys []string `envconfig:"api_keys" required:"true"`
}

var configuration Configuration

// GetConfiguration ...
func GetConfiguration() *Configuration {
	return &configuration
}

func init() {
	err := envconfig.Process("", &configuration)
	if err != nil {
		log.Fatal(err)
	}
}
