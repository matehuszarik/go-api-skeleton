package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Configuration ..
type Configuration struct {
	Port int `envconfig:"port" default:"80"`
}

var configuration Configuration

// GetConfiguration ...
func GetConfiguration() *Configuration {
	return &configuration
}

func init() {
	err := envconfig.Process("", &configuration)
	if err != nil {
		log.Fatal("Can not parse environment variables", err)
	}
}
