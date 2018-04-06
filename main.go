package main

import (
	"github.com/matehuszarik/go-api-skeleton/internal"
	"github.com/matehuszarik/go-api-skeleton/internal/config"
)

func main() {
	c := config.GetConfiguration()

	s := internal.NewServer(c.APIKeys)
	s.Serve(c.Port)
}
