package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/auth/pb"
)

type AuthMiddleWareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleWare(svc *ServiceClient) AuthMiddleWareConfig {
	return AuthMiddleWareConfig{svc}
}

func (c *AuthMiddleWareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")
	fmt.Println("token is ", token)

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})
	fmt.Println(err, res)

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.Set("userId", res.UserId)
	ctx.Next()
}
