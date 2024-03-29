package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/auth/routes"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/config"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/middleware"
)

func RegisterRoutes(r *gin.Engine, c config.Config) *ServiceClient {

	svc := &ServiceClient{
		Client: InitServiceClient(&c),
	}
	authMiddleware := middleware.InitAuthMiddleWare(c)
	adminAuthMiddleware := authMiddleware.AdminAuthRequired // Take the address here
	userAuthMiddleware := authMiddleware.UserAuthRequired

	//roots accessible for user

	authRoutes := r.Group("/auth")

	authRoutes.POST("/register", svc.Register)
	authRoutes.POST("/login", svc.Login)

	//middleware for user

	authRoutes.Use(userAuthMiddleware)
	authRoutes.PATCH("/resetPassword", userAuthMiddleware, svc.ResetPassword)
	authRoutes.GET("/fetchUser", userAuthMiddleware, svc.FetchShortDetails)

	//roots accesible for admin

	adminRoutes := r.Group("/admin")

	adminRoutes.POST("/adminsignup", svc.AdminSignup)
	adminRoutes.POST("adminlogin", svc.AdminLogin)

	//middleware for admin

	adminRoutes.Use(adminAuthMiddleware)
	authRoutes.GET("/checkBlock", adminAuthMiddleware, svc.CheckUserBlocked)
	authRoutes.PATCH("/block", adminAuthMiddleware, svc.BlockUser)
	authRoutes.PATCH("/unblock", adminAuthMiddleware, svc.UnblockUser)

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
func (svc *ServiceClient) ResetPassword(ctx *gin.Context) {
	routes.ResetPassword(ctx, svc.Client)
}
func (svc *ServiceClient) ForgotPassWord(ctx *gin.Context) {
	routes.ForgotPassWord(ctx, svc.Client)
}
func (svc *ServiceClient) CheckUserBlocked(ctx *gin.Context) {
	routes.CheckUserBlocked(ctx, svc.Client)
}
func (svc *ServiceClient) BlockUser(ctx *gin.Context) {
	routes.BlockUser(ctx, svc.Client)
}
func (svc *ServiceClient) UnblockUser(ctx *gin.Context) {
	routes.UnblockUser(ctx, svc.Client)
}
func (svc *ServiceClient) FetchShortDetails(ctx *gin.Context) {
	routes.FetchShortDetails(ctx, svc.Client)
}
