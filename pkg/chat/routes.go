package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/auth"
	routsC "github.com/sgokul961/echo-hub-api-gateway/pkg/chat/routs"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c config.Config, authSvc *auth.ServiceClient) {

	svc := &ServiceClient{
		Client: InitServiceClient(&c),
	}

	authMiddleware := auth.InitAuthMiddleWare(authSvc, c)
	userAuthMiddleware := authMiddleware.UserAuthRequired

	chatRoutes := r.Group("/chat")
	chatRoutes.GET("/ws", userAuthMiddleware, svc.Chat)
	chatRoutes.GET("", userAuthMiddleware, svc.Getchats)
	chatRoutes.GET("/:chatId/message", userAuthMiddleware, svc.GetMessages)

}
func (svc *ServiceClient) Chat(ctx *gin.Context) {

	routsC.Chat(ctx, svc.Client)

}
func (svc *ServiceClient) Getchats(ctx *gin.Context) {
	routsC.Getchats(ctx, svc.Client)
}
func (svc *ServiceClient) GetMessages(ctx *gin.Context) {
	routsC.GetMessages(ctx, svc.Client)
}
