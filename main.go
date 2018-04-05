package main

import (
	"github.com/matehuszarik/go-api-skeleton/internal"
	"github.com/matehuszarik/go-api-skeleton/internal/config"
)

func main() {
	s := internal.NewServer()
	s.Serve(config.GetConfiguration().Port)
}
