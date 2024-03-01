package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/models"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/post/pb"
)

func FollowUser(ctx *gin.Context, p pb.PostServiceClient) {
	userid, ok := ctx.Get("userId")
	if !ok || userid == nil {
		ctx.JSON(401, gin.H{"error": "userId not found in context or is nil"})
		return
	}
	followId := ctx.Param("follow_id")

	followerIDInt, err := strconv.ParseInt(followId, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid followId format"})
		return
	}
	fmt.Println("follower ", followerIDInt)

	response, err := p.FollowUser(ctx, &pb.FollowUserRequest{
		FollowUserId:   userid.(int64),
		FollowerUserId: followerIDInt,
	})

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	successRes := models.MakeResponse(http.StatusOK, "successfully followed user", response, nil)
	ctx.JSON(http.StatusOK, successRes)

}
func UnfollowUser(ctx *gin.Context, p pb.PostServiceClient) {

	following_user_id, ok := ctx.Get("userId")
	if !ok || following_user_id == nil {
		ctx.JSON(401, gin.H{"error": "following userId not found in context or is nil"})
		return
	}
	followId := ctx.Param("unfollow_id")

	followerIDInt, err := strconv.ParseInt(followId, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid followId format"})
		return
	}
	res, err := p.UnfollowUser(ctx, &pb.UnfollowUserRequest{
		FollowUserId:   following_user_id.(int64),
		FollowerUserId: followerIDInt,
	})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	successRes := models.MakeResponse(http.StatusOK, "successfully unfollowed user", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}
