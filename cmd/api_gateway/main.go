package main

import (
	"log"

	_ "github.com/sgokul961/echo-hub-api-gateway/cmd/api_gateway/docs"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/config"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/server"
)

//	@title			ECHO-HUB-SOCIAL-MEADIA
//	@version		1.0
//	@description	WELCOME TO ECHO-HUB
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize server
	srv := server.NewServer(&cfg)

	// Start server
	if err := srv.Run(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
