package chat

import (
	"fmt"

	"github.com/sgokul961/echo-hub-api-gateway/pkg/chat/pb"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ChatServiceClient
}

func InitServiceClient(c *config.Config) pb.ChatServiceClient {
	cc, err := grpc.Dial(c.ChatHubUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("coudnt connect:", err)
	}

	return pb.NewChatServiceClient(cc)

}
