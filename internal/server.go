package internal

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matehuszarik/go-api-skeleton/internal/handlers"
)

// Server ...
type Server struct {
	router *gin.Engine
}

// NewServer ...
func NewServer() Server {
	s := Server{
		router: gin.New(),
	}

	s.router.Use(
		gin.Logger(),
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
func (s Server) Serve(port int) {
	s.router.Run(fmt.Sprint(":", port))
}
