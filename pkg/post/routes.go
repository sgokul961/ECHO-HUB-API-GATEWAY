package post

import (
	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/auth"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/config"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/post/routes"
)

func RegisterRoutes(r *gin.Engine, c config.Config, authSvc *auth.ServiceClient) {

	svc := &ServiceClient{
		Client: InitServiceClient(&c),
	}

	authMiddleware := auth.InitAuthMiddleWare(authSvc, c)
	userAuthMiddleware := authMiddleware.UserAuthRequired

	postrouts := r.Group("/post")
	//r.Use(userAuthMiddleware)
	postrouts.POST("/follow/:follow_id", userAuthMiddleware, svc.FollowUser)

}

func (svc *ServiceClient) FollowUser(ctx *gin.Context) {
	routes.FollowUser(ctx, svc.Client)
}
