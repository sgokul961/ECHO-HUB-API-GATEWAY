package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/config"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/response"
)

type jwtClaims struct {
	jwt.StandardClaims
	Id   int64
	Role string
}
type AuthMiddleWareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleWare(svc *ServiceClient) *AuthMiddleWareConfig {
	return &AuthMiddleWareConfig{svc: svc}
}

func (c *AuthMiddleWareConfig) AdminAuthRequired(ctx *gin.Context) {

	authorization := ctx.Request.Header.Get("Authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	validastring, err := validateToken(token[1])

	if err != nil || !validastring.Valid {
		errRes := response.MakeResponse(http.StatusUnauthorized, "not authorised for this action", nil, err)
		ctx.JSON(http.StatusUnauthorized, errRes)
		ctx.Abort()
		return
	}
	claims, ok := validastring.Claims.(*jwtClaims)

	if !ok || claims == nil {
		errRes := response.MakeResponse(http.StatusUnauthorized, "not authorised for this action", nil, fmt.Errorf("unable to retrieve claims"))
		ctx.JSON(http.StatusUnauthorized, errRes)
		ctx.Abort()
		return
	}
	if claims.Role == "user" {
		errRes := response.MakeResponse(http.StatusUnauthorized, "users are not authorised for this action", nil, err)
		ctx.JSON(http.StatusUnauthorized, errRes)
		ctx.Abort()
		return

	}
	userID := claims.Id
	role := claims.Role
	ctx.Set("userId", userID)
	ctx.Set("role", role)

	// Proceed with the authenticated request.
	ctx.Next()

}

func validateToken(tokenString string) (*jwt.Token, error) {
	var con config.Config
	token, err := jwt.ParseWithClaims(tokenString, jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(con.JWTSecretKey), nil
	})

	return token, err
}
func (c *AuthMiddleWareConfig) UserAuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token := strings.Split(authorization, "Bearer ")

	validastring, err := validateToken(token[1])

	if err != nil || !validastring.Valid {
		errRes := response.MakeResponse(http.StatusUnauthorized, "not authorised for this action", nil, err)
		ctx.JSON(http.StatusUnauthorized, errRes)
		ctx.Abort()
		return
	}
	claims, ok := validastring.Claims.(*jwtClaims)

	if !ok || claims == nil {
		errRes := response.MakeResponse(http.StatusUnauthorized, "not authorised for this action", nil, fmt.Errorf("unable to retrieve claims"))
		ctx.JSON(http.StatusUnauthorized, errRes)
		ctx.Abort()
		return
	}
	userId := claims.Id
	role := claims.Role

	ctx.Set("userId", userId)
	ctx.Set("role", role)

	// Proceed with the authenticated request.
	ctx.Next()

}
