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
	con config.Config
}

func InitAuthMiddleWare(svc *ServiceClient, con config.Config) *AuthMiddleWareConfig {
	return &AuthMiddleWareConfig{svc: svc, con: con}
}

func (c *AuthMiddleWareConfig) AdminAuthRequired(ctx *gin.Context) {

	authorization := ctx.Request.Header.Get("Authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")
	if len(token) != 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	validastring, err := c.validateToken(token[1])

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

func (c *AuthMiddleWareConfig) validateToken(tokenString string) (*jwt.Token, error) {
	//var con config.Config
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		fmt.Println("jwt key", c.con.JWTSecretKey)
		return []byte(c.con.JWTSecretKey), nil
	})
	if err != nil {
		// There was an error during token validation
		fmt.Println("Error validating token:", err)
	}

	if claims, ok := token.Claims.(*jwtClaims); ok && token.Valid {
		fmt.Printf("Token ID: %d\n", claims.Id)
		fmt.Printf("Token Role: %s\n", claims.Role)
		// Add more logging or inspection of claims as needed
	} else {
		fmt.Println("Invalid token claims")
	}

	return token, err
}
func (c *AuthMiddleWareConfig) UserAuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token := strings.Split(authorization, "Bearer ")
	fmt.Println("token", token)
	if len(token) != 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	validastring, err := c.validateToken(token[1])
	fmt.Println("validtokenor not", validastring)

	if err != nil || !validastring.Valid {
		errRes := response.MakeResponse(http.StatusUnauthorized, "not authorised for this action", nil, err)
		ctx.JSON(http.StatusUnauthorized, errRes)
		ctx.Abort()
		return
	}
	claims, ok := validastring.Claims.(*jwtClaims)

	fmt.Println("calims role", claims.Role)

	if !ok || claims == nil {
		errRes := response.MakeResponse(http.StatusUnauthorized, "claims nil ,not authorised for this action", nil, fmt.Errorf("unable to retrieve claims"))
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

// func (c *AuthMiddleWareConfig) validateToken(tokenString string) (*jwtClaims, error) {
// 	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
// 		}
// 		return []byte(c.con.JWTSecretKey), nil
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("error parsing token: %v", err)
// 	}

// 	// Check if token is valid and retrieve claims
// 	if claims, ok := token.Claims.(*jwtClaims); ok && token.Valid {
// 		return claims, nil
// 	}

// 	// Invalid token or claims
// 	return nil, fmt.Errorf("invalid token or claims")
// }
