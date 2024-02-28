package post

import (
	"fmt"

	"github.com/sgokul961/echo-hub-api-gateway/pkg/config"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/post/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.PostServiceClient
}

func InitServiceClient(c *config.Config) pb.PostServiceClient {
	cc, err := grpc.Dial(c.PostHubUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("coudnt connect:", err)
	}
	return pb.NewPostServiceClient(cc)

}
