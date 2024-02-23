package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/auth/pb"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/models"
)

func Login(ctx *gin.Context, p pb.AuthServiceClient) {

	var userLogin pb.LoginRequest
	err := ctx.BindJSON(&userLogin)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error parsing request body", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	// Call the AuthServiceClient to handle the login process
	res, err := p.Login(ctx, &userLogin)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error connecting auth service", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	fmt.Println("res", res)

	successRes := models.MakeResponse(http.StatusOK, "user successfully logged in", res, nil)

	ctx.JSON(http.StatusOK, successRes)

}
