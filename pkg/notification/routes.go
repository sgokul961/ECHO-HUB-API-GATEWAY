package notification

import (
	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/auth"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/config"
	notification_routes "github.com/sgokul961/echo-hub-api-gateway/pkg/notification/routes"
)

func RegisterRoutes(r *gin.Engine, c config.Config, authSvc *auth.ServiceClient) {

	svc := &ServiceClient{
		Client: InitServiceClient(&c),
	}

	authMiddleware := auth.InitAuthMiddleWare(authSvc, c)
	userAuthMiddleware := authMiddleware.UserAuthRequired

	notifyRouts := r.Group("/notify")
	notifyRouts.POST("/commentNotification", userAuthMiddleware, svc.SendCommentedNotification)

}

func (svc *ServiceClient) SendCommentedNotification(ctx *gin.Context) {

	notification_routes.SendCommentedNotification(ctx, svc.Client)

}
