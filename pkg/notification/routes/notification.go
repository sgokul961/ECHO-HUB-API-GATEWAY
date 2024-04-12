package routes

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/models"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/notification/pb"
)

func SendCommentedNotification(ctx *gin.Context, p pb.NotificationServiceClient) {

}
func SendLikeNotification(ctx *gin.Context, p pb.NotificationServiceClient) {
	var notification pb.LikeNotification
	userID, ok := ctx.Get("userId")
	if !ok || userID == nil {
		errRes := models.MakeResponse(http.StatusUnauthorized, "user id is not valid", nil, errors.New("user id is nil"))
		ctx.JSON(http.StatusUnauthorized, errRes)
		return
	}

	PostId := ctx.Param("postId")

	postIdint, err := strconv.ParseInt(PostId, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid post id format"})
		return
	}
	notification.PostId = postIdint

	res, err := p.SendLikeNotification(ctx, &notification)

	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error in connecting with  notification service", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	successRes := models.MakeResponse(http.StatusOK, "successfully got all notifications  ", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}

//stream notification messages messages

func ConsumeKafkaMessages(ctx *gin.Context, p pb.NotificationServiceClient) {
	// Call the gRPC method to start streaming Kafka messages
	stream, err := p.ConsumeKafkaMessages(ctx, &pb.Empty{})
	if err != nil {
		errRes := models.MakeResponse(http.StatusInternalServerError, "error while receiving Kafka message from stream", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, errRes)
		return
	}

	// Stream Kafka messages to the client
	ctx.Stream(func(w io.Writer) bool {
		// Receive next Kafka message from the stream
		message, err := stream.Recv()
		if err == io.EOF {
			// End of stream
			return false
		}
		if err != nil {
			// Handle error
			errRes := models.MakeResponse(http.StatusInternalServerError, "error while receiving Kafka message from stream", nil, err.Error())
			ctx.JSON(http.StatusInternalServerError, errRes)
			return false
		}

		// Process the Kafka message as needed (e.g., log, store in database, etc.)
		fmt.Printf("Received Kafka message: %+v\n", message)

		// Send the Kafka message to the client
		ctx.SSEvent("message", message)

		// Continue streaming
		return true
	})
	// Close the stream
	stream.CloseSend()
}

// func ConsumeKafkaMessages(ctx *gin.Context, p pb.NotificationServiceClient) {
// 	// Call the gRPC method to start streaming Kafka messages
// 	stream, err := p.ConsumeKafkaMessages(ctx, &pb.Empty{})
// 	if err != nil {
// 		errRes := models.MakeResponse(http.StatusInternalServerError, "error while receiving Kafka message from stream", nil, err.Error())
// 		ctx.JSON(http.StatusInternalServerError, errRes)
// 		return
// 	}

// 	// Variable to track if a message has been received
// 	messageReceived := false

// 	// Stream Kafka messages to the client
// 	ctx.Stream(func(w io.Writer) bool {
// 		if messageReceived {
// 			// If a message has already been received, close the stream
// 			return false
// 		}

// 		// Receive next Kafka message from the stream
// 		message, err := stream.Recv()
// 		if err == io.EOF {
// 			// End of stream
// 			return false
// 		}
// 		if err != nil {
// 			// Handle error
// 			errRes := models.MakeResponse(http.StatusInternalServerError, "error while receiving Kafka message from stream", nil, err.Error())
// 			ctx.JSON(http.StatusInternalServerError, errRes)
// 			return false
// 		}

// 		// Process the Kafka message as needed (e.g., log, store in database, etc.)
// 		fmt.Printf("Received Kafka message: %+v\n", message)

// 		// Send the Kafka message to the client
// 		ctx.SSEvent("message", message)

// 		// Set messageReceived to true to indicate that a message has been received
// 		messageReceived = true

// 		// Continue streaming
// 		return true
// 	})

// Close the stream
//
//		stream.CloseSend()
//	}
func ConsumeKafkaCommentMessages(ctx *gin.Context, p pb.NotificationServiceClient) {
	var getUserId pb.ConsumeKafkaCommentMessagesRequest

	userID, ok := ctx.Get("userId")

	fmt.Println("user id ", userID)
	if !ok || userID == nil {
		errRes := models.MakeResponse(http.StatusUnauthorized, "user id is not valid", nil, errors.New("user id is nil"))
		ctx.JSON(http.StatusUnauthorized, errRes)
		return
	}
	getUserId.UserId = userID.(int64)

	res, err := p.ConsumeKafkaCommentMessages(ctx, &getUserId)

	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error connecting notification service", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	successRes := models.MakeResponse(http.StatusOK, "New Comment Message", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}
func ConsumeKafkaLikeMessages(ctx *gin.Context, p pb.NotificationServiceClient) {

	var getUserId pb.ConsumeKafkaLikeMessagesRequest

	userID, ok := ctx.Get("userId")

	fmt.Println("user id ", userID)
	if !ok || userID == nil {
		errRes := models.MakeResponse(http.StatusUnauthorized, "user id is not valid", nil, errors.New("user id is nil"))
		ctx.JSON(http.StatusUnauthorized, errRes)
		return
	}
	getUserId.UserId = userID.(int64)

	res, err := p.ConsumeKafkaLikeMessages(ctx, &getUserId)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error connecting notification service", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	successRes := models.MakeResponse(http.StatusOK, "New like Added", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}
