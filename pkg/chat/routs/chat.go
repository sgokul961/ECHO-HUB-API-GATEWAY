package routsC

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/chat/pb"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/models"
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

func Chat(ctx *gin.Context, p pb.ChatServiceClient) {

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "data is not in required format", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	userid, ok := ctx.Get("userId")

	if !ok || userid == nil {
		ctx.JSON(401, gin.H{"error": "userId not found in context or is nil"})
		return
	}

	userIdInt64 := userid.(int64)
	//==================================================================================//getting user id from context

	chatID := ctx.Query("chatID")
	if chatID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "chatID parameter is missing"})
		return
	}

	objectChatId, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		fmt.Println("Error converting chatID to ObjectID:", err)
		return
	}
	//=================================================================================//getting  chat id from query parameater

	connection[conn] = &client{ChatId: objectChatId, UserId: userIdInt64}
	user[userIdInt64] = conn
	go func() {

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			userId := connection[conn].UserId
			chatID := connection[conn].ChatId

			_, err = p.SaveMessage(ctx, &pb.SaveMessageRequest{
				ChatId:   objectChatId.Hex(),
				UserId:   userIdInt64,
				Messages: string(msg),
			})
			if err != nil {
				fmt.Println("Error saving message:", err)
				break
			}

			conn.WriteMessage(websocket.TextMessage, msg)

			recipient, err := p.FetchRecipient(ctx, &pb.FetchRecipientRequest{
				ChatId: chatID.Hex(),
				UserId: userId,
			})
			if err != nil {
				fmt.Println("Error fetching recipient:", err)
				break
			}
			recipientID := recipient.GetRecipient()

			if value, ok := user[recipientID]; ok {
				err = value.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					delete(connection, value)
					delete(user, recipientID)
				}
			}
		}
	}()
}
func Getchats(ctx *gin.Context, p pb.ChatServiceClient) {

	var GetChat pb.GetchatsRequest

	userID, ok := ctx.Get("userId")
	if !ok || userID == nil {
		errRes := models.MakeResponse(http.StatusUnauthorized, "user id is not valid", nil, errors.New("user id is nil"))
		ctx.JSON(http.StatusUnauthorized, errRes)
		return
	}

	GetChat.UserId = userID.(int64)

	res, err := p.Getchats(ctx, &GetChat)

	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error in connecting with  chat service ", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	successRes := models.MakeResponse(http.StatusOK, "successfully got  all chats ", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}
func GetMessages(ctx *gin.Context, p pb.ChatServiceClient) {

	var GetMessage pb.GetMessagesRequest

	chatID := ctx.Param("chatId")
	fmt.Println("chat id is ", chatID)
	if chatID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "chatID parameter is missing"})
		return
	}
	GetMessage.ChatId = chatID

	res, err := p.GetMessages(ctx, &GetMessage)

	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error in connecting with  chat service ", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	successRes := models.MakeResponse(http.StatusOK, "successfully got  all messages ", res.GetMessages(), nil)
	ctx.JSON(http.StatusOK, successRes)

}
