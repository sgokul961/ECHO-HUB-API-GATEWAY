package auth

import (
	"fmt"

	"github.com/sgokul961/echo-hub-api-gateway/pkg/auth/pb"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {

	//if call fails chech for error here grpc communication

	cc, err := grpc.Dial(c.AuthHubUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("coudnt connect:", err)
	}
	return pb.NewAuthServiceClient(cc)
}
