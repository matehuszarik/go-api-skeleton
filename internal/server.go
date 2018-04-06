package internal

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matehuszarik/go-api-skeleton/internal/handlers"
	"github.com/matehuszarik/go-api-skeleton/internal/middlewares"
)

// Server ...
type Server struct {
	router *gin.Engine
}

// NewServer ...
func NewServer(apiKeys []string) Server {
	s := Server{
		router: gin.New(),
	}

	s.router.Use(
		gin.Logger(),
		middlewares.APIKey(apiKeys),
		gin.Recovery(),
	)

	s.setupHandlers()

	return s
}

func (s Server) setupHandlers() {
	api := s.router.Group("/api")

	api.Use()
	{
		api.GET("/healthcheck", handlers.HealthcheckHandler)
	}
}

// Serve ...
func (s Server) Serve(port int) error {
	return s.router.Run(fmt.Sprint(":", port))
}

// ServeTLS ...
func (s Server) ServeTLS(port int, certFile string, keyFile string) error {
	return s.router.RunTLS(fmt.Sprint(":", port), "", "")
}
