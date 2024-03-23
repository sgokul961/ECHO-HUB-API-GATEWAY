package notification

import (
	"fmt"

	"github.com/sgokul961/echo-hub-api-gateway/pkg/config"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/notification/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.NotificationServiceClient
}

func InitServiceClient(c *config.Config) pb.NotificationServiceClient {
	cc, err := grpc.Dial(c.NotificationHubUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("coudnt connect:", err)
	}

	return pb.NewNotificationServiceClient(cc)

}
