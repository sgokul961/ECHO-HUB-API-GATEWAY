package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/auth"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/config"
)

type Server struct {
	router *gin.Engine
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	r := gin.Default()

	// Register routes for authentication service
	auth.RegisterRoutes(r, *cfg)

	return &Server{
		router: r,
		config: cfg,
	}
}

// Run starts the API gateway server
func (s *Server) Run() error {
	err := s.router.Run(":" + s.config.Port)
	if err != nil {
		return err
	}
	return nil
}
