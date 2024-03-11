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
	postrouts.DELETE("/unfollow/:unfollow_id", userAuthMiddleware, svc.UnfollowUser)
	postrouts.POST("/upload_post", userAuthMiddleware, svc.UploadPost)
	postrouts.DELETE("/deletePost/:post_id", userAuthMiddleware, svc.DeletePost)
	postrouts.POST("/like/:postId", userAuthMiddleware, svc.LikePost)
	postrouts.DELETE("/dislike/:postId", userAuthMiddleware, svc.DislikePost)
	postrouts.POST("/comment/:post_id", userAuthMiddleware, svc.CommentPost)
	postrouts.GET("/getcomment", userAuthMiddleware, svc.GetComments)
	postrouts.DELETE("/:post_id", userAuthMiddleware, svc.DeleteComments)

}

func (svc *ServiceClient) FollowUser(ctx *gin.Context) {
	routes.FollowUser(ctx, svc.Client)
}
func (svc *ServiceClient) UnfollowUser(ctx *gin.Context) {
	routes.UnfollowUser(ctx, svc.Client)
}
func (svc *ServiceClient) UploadPost(ctx *gin.Context) {
	routes.UploadPost(ctx, svc.Client)
}
func (svc *ServiceClient) DeletePost(ctx *gin.Context) {
	routes.DeletePost(ctx, svc.Client)
}
func (svc *ServiceClient) LikePost(ctx *gin.Context) {
	routes.LikePost(ctx, svc.Client)
}
func (svc *ServiceClient) DislikePost(ctx *gin.Context) {
	routes.DislikePost(ctx, svc.Client)
}
func (svc *ServiceClient) CommentPost(ctx *gin.Context) {
	routes.CommentPost(ctx, svc.Client)
}
func (svc *ServiceClient) GetComments(ctx *gin.Context) {
	routes.GetComments(ctx, svc.Client)
}
func (svc *ServiceClient) DeleteComments(ctx *gin.Context) {
	routes.DeleteComments(ctx, svc.Client)
}
