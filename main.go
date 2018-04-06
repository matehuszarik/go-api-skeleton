package main

import (
	"github.com/matehuszarik/go-api-skeleton/internal"
	"github.com/matehuszarik/go-api-skeleton/internal/config"
)

func main() {
	s := internal.NewServer(config.GetConfiguration().APIKeys)
	s.Serve(config.GetConfiguration().Port)
}
