package routes

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/helper"
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
func UploadPost(ctx *gin.Context, p pb.PostServiceClient) {
	userId, ok := ctx.Get("userId")
	if !ok || userId == nil {
		ctx.JSON(401, gin.H{"error": "following userId not found in context or is nil"})
		return
	}
	var req pb.UploadPostRequest
	req.UserId = userId.(int64)
	req.Content = ctx.Request.FormValue("content")

	files := ctx.Request.MultipartForm.File["media_urls"]
	var mediaUrls []string

	// Iterate over each file
	for _, fileHeader := range files {
		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		defer file.Close()

		// Upload the file to S3 and get the URL
		picture, err := helper.AddImageToS3(fileHeader)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Add the URL to the mediaUrls slice
		mediaUrls = append(mediaUrls, picture)
	}

	// Set the mediaUrls field in the request
	req.MediaUrls = mediaUrls

	res, err := p.UploadPost(ctx, &req)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error in connecting with service", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	succRes := models.MakeResponse(http.StatusOK, "successfully uploaded  post", res, nil)

	ctx.JSON(http.StatusOK, succRes)

}
func DeletePost(ctx *gin.Context, p pb.PostServiceClient) {

	var req pb.DeletePostRequest

	user_id, ok := ctx.Get("userId")
	if !ok || user_id == nil {
		ctx.JSON(401, gin.H{"error": " userId not found in context or is nil"})
		return
	}

	req.UserId = user_id.(int64)
	fmt.Println("req.user_id", req.UserId)

	post_id := ctx.Param("post_id")

	postIdint, err := strconv.ParseInt(post_id, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid post id format"})
		return
	}
	req.PostId = postIdint

	res, err := p.DeletePost(ctx, &req)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error in connecting with delete postservice", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	successRes := models.MakeResponse(http.StatusOK, "successfully deleted post", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}
func LikePost(ctx *gin.Context, p pb.PostServiceClient) {

	var like pb.LikePostRequest
	user_id, ok := ctx.Get("userId")

	if !ok || user_id == nil {
		errRes := models.MakeResponse(http.StatusUnauthorized, "user id is not valid ", nil, errors.New("user id is nil"))
		ctx.JSON(http.StatusBadGateway, errRes)
		return

	}
	like.UserId = user_id.(int64)

	PostId := ctx.Param("postId")
	postIdint, err := strconv.ParseInt(PostId, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid post id format"})
		return
	}
	like.PostId = postIdint

	res, err := p.LikePost(ctx, &like)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error in connecting with like postservice", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	successRes := models.MakeResponse(http.StatusOK, "successfully liked post", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}
func DislikePost(ctx *gin.Context, p pb.PostServiceClient) {
	var dislike pb.DislikePostRequest
	user_id, ok := ctx.Get("userId")

	if !ok || user_id == nil {
		errRes := models.MakeResponse(http.StatusUnauthorized, "user id is not valid ", nil, errors.New("user id is nil"))
		ctx.JSON(http.StatusBadGateway, errRes)
		return

	}
	dislike.UserId = user_id.(int64)

	PostId := ctx.Param("postId")
	postIdint, err := strconv.ParseInt(PostId, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid post id format"})
		return
	}
	dislike.PostId = postIdint
	res, err := p.DislikePost(ctx, &dislike)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error in connecting with like postservice", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}

	successRes := models.MakeResponse(http.StatusOK, "successfully disliked post", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}
func CommentPost(ctx *gin.Context, p pb.PostServiceClient) {
	// Parse user ID from context
	userID, ok := ctx.Get("userId")
	if !ok || userID == nil {
		errRes := models.MakeResponse(http.StatusUnauthorized, "user id is not valid", nil, errors.New("user id is nil"))
		ctx.JSON(http.StatusUnauthorized, errRes)
		return
	}
	fmt.Println("user id", userID)

	// Parse post ID from URL parameter
	postIDStr := ctx.Param("post_id")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post id format"})
		return
	}
	fmt.Println("post id is", postID)

	var comment pb.CommentPostRequest

	// Assign user ID, post ID, and content to the comment

	comment.UserId = userID.(int64)
	comment.PostId = postID
	fmt.Println("postid", postID)
	comment.Content = ctx.Request.FormValue("content")
	fmt.Println("content", comment.Content)

	// Call the gRPC method to send the comment to the post service
	res, err := p.CommentPost(ctx, &comment)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error connecting to the post service", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}

	// Respond with success message
	successRes := models.MakeResponse(http.StatusOK, "comment successfully added", res, nil)
	ctx.JSON(http.StatusOK, successRes)
}
func GetComments(ctx *gin.Context, p pb.PostServiceClient) {

	var GetComment pb.GetCommentsRequest

	err := ctx.BindJSON(&GetComment)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error parsing request body,email ", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}

	res, err := p.GetComments(ctx, &GetComment)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error connecting post service", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	successRes := models.MakeResponse(http.StatusOK, "Successfully got all comments", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}
func DeleteComments(ctx *gin.Context, p pb.PostServiceClient) {

	userID, ok := ctx.Get("userId")

	if !ok || userID == nil {
		errRes := models.MakeResponse(http.StatusUnauthorized, "user id is not valid", nil, errors.New("user id is nil"))
		ctx.JSON(http.StatusUnauthorized, errRes)
		return
	}

	// Parse post ID from URL parameter
	postIDStr := ctx.Param("post_id")

	postID, err := strconv.ParseInt(postIDStr, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post id format"})
		return
	}

	var deleteComment pb.DeleteCommentRequest

	deleteComment.UserId = userID.(int64)

	deleteComment.PostId = postID

	if err := ctx.BindJSON(&deleteComment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	res, err := p.DeleteComments(ctx, &deleteComment)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error connecting post service", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}

	successRes := models.MakeResponse(http.StatusOK, "successfully deleted  comment", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}
func GetUserId(ctx *gin.Context, p pb.PostServiceClient) {

	var getUserIdOfPost pb.GetUserIdRequest

	err := ctx.BindJSON(&getUserIdOfPost)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error parsing request body ", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}

	res, err := p.GetUserId(ctx, &getUserIdOfPost)

	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error connecting post service", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	successRes := models.MakeResponse(http.StatusOK, "Successfully got userId", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}
