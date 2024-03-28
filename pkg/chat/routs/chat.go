package routsC

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/chat/pb"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type client struct {
	ChatId primitive.ObjectID
	UserId int64
}

var (
	connection = make(map[*websocket.Conn]*client)
	user       = make(map[int64]*websocket.Conn)
)

// func Chat(ctx *gin.Context, p pb.ChatServiceClient) {
// 	userid, ok := ctx.Get("userId")
// 	if !ok || userid == nil {
// 		ctx.JSON(401, gin.H{"error": "userId not found in context or is nil"})
// 		return
// 	}

// 	chatID := ctx.Param("chatId")

// 	// Assuming you have a context to pass to the gRPC call, if not, use `context.Background()`
// 	res, err := p.Chat(ctx, &pb.SendMessageRequest{
// 		UserId: userid.(int64),
// 		ChatId: chatID, // chatID is already a string
// 	})
// 	if err != nil {
// 		ctx.JSON(500, gin.H{"error": "Failed to send message"})
// 		return
// 	}
// 	successRes := models.MakeResponse(http.StatusOK, "successfully unfollowed user", res, nil)
// 	ctx.JSON(http.StatusOK, successRes)
// }

// func Chat(ctx *gin.Context, p pb.ChatServiceClient) {
// 	//conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
// 	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

// 	if err != nil {
// 		errRes := response.MakeResponse(http.StatusBadRequest, "data is not in required format", nil, err.Error())
// 		ctx.JSON(http.StatusBadRequest, errRes)
// 		return
// 	}
// 	userid, ok := ctx.Get("userId")
// 	if !ok || userid == nil {
// 		ctx.JSON(401, gin.H{"error": "userId not found in context or is nil"})
// 		return
// 	}
// 	userIdInt64 := userid.(int64)
// 	chat_id, err := primitive.ObjectIDFromHex(ctx.Param("chatId"))
// 	if err != nil {
// 		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
// 		ctx.JSON(http.StatusBadRequest, errRes)
// 		return
// 	}

// 	connection[conn] = &client{ChatId: chat_id, UserId: userIdInt64}
// 	user[userIdInt64] = conn
// 	go func() {

// 		for {
// 			_, msg, err := conn.ReadMessage()
// 			if err != nil {
// 				break
// 			}
// 			userId := connection[conn].UserId
// 			chatID := connection[conn].ChatId

// 			_, err = p.SaveMessage(ctx, &pb.SaveMessageRequest{
// 				ChatId:   chatID.Hex(),
// 				UserId:   userIdInt64,
// 				Messages: string(msg),
// 			})
// 			if err != nil {
// 				fmt.Println("Error saving message:", err)
// 				break
// 			}

// 			conn.WriteMessage(websocket.TextMessage, msg)

// 			recipient, err := p.FetchRecipient(ctx, &pb.FetchRecipientRequest{
// 				ChatId: chatID.Hex(),
// 				UserId: userId,
// 			})
// 			if err != nil {
// 				fmt.Println("Error fetching recipient:", err)
// 				break
// 			}
// 			recipientID := recipient.GetRecipient()

// 			if value, ok := user[recipientID]; ok {
// 				err = value.WriteMessage(websocket.TextMessage, msg)
// 				if err != nil {
// 					delete(connection, value)
// 					delete(user, recipientID)
// 				}
// 			}
// 		}
// 	}()

// }

func Chat(ctx *gin.Context, p pb.ChatServiceClient) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "data is not in required format", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}
	fmt.Println("WebSocket connection established.")

	userid, ok := ctx.Get("userId")
	if !ok || userid == nil {
		ctx.JSON(401, gin.H{"error": "userId not found in context or is nil"})
		return
	}
	userIdInt64 := userid.(int64)

	fmt.Println("useridint:", userIdInt64)
	chat_id, err := primitive.ObjectIDFromHex(ctx.Param("chatId"))

	fmt.Println("chatid from hex :", chat_id)
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	connection[conn] = &client{ChatId: chat_id, UserId: userIdInt64}
	user[userIdInt64] = conn
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			fmt.Println("message is ", msg)
			if err != nil {
				break
			}
			userId := connection[conn].UserId
			chatID := connection[conn].ChatId

			_, err = p.SaveMessage(ctx, &pb.SaveMessageRequest{
				ChatId:   chatID.Hex(),
				UserId:   userIdInt64,
				Messages: string(msg),
			})
			if err != nil {
				fmt.Println("Error saving message:", err)
				break
			}

			conn.WriteMessage(websocket.TextMessage, msg) // Changed from ctx.Writer to conn

			recipient, err := p.FetchRecipient(ctx, &pb.FetchRecipientRequest{
				ChatId: chatID.Hex(),
				UserId: userId,
			})
			fmt.Println("recipient ", recipient)
			if err != nil {
				fmt.Println("Error fetching recipient:", err)
				break
			}
			recipientID := recipient.GetRecipient()

			if value, ok := user[recipientID]; ok {
				err = value.WriteMessage(websocket.TextMessage, msg) // Changed from ctx.Writer to value
				if err != nil {
					delete(connection, value)
					delete(user, recipientID)
				}
			}
		}
	}()
}
