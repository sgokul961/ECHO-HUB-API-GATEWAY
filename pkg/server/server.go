package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/auth"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/config"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/post"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Server struct {
	router *gin.Engine
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Register routes for authentication service
	authSvc := auth.RegisterRoutes(r, *cfg)
	post.RegisterRoutes(r, *cfg, authSvc)

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
