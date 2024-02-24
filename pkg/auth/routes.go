package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/auth/routes"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	authMiddleware := InitAuthMiddleWare(svc)
	adminAuthMiddleware := authMiddleware.AdminAuthRequired // Take the address here
	userAuthMiddleware := authMiddleware.UserAuthRequired

	//roots accessible for user

	authRoutes := r.Group("/auth")

	//middleware for user
	authRoutes.Use(userAuthMiddleware)
	authRoutes.POST("/register", svc.Register)
	authRoutes.POST("/login", svc.Login)

	//roots accesible for admin

	adminRoutes := r.Group("/admin")

	//middleware for admin

	adminRoutes.Use(adminAuthMiddleware)

	adminRoutes.POST("/adminsignup", svc.AdminSignup)
	adminRoutes.POST("adminlogin", svc.AdminLogin)

	return svc

}
func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}
func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
func (svc *ServiceClient) AdminLogin(ctx *gin.Context) {
	routes.AdminLogin(ctx, svc.Client)
}
func (svc *ServiceClient) AdminSignup(ctx *gin.Context) {
	routes.AdminSignup(ctx, svc.Client)
}
