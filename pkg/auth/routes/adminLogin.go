package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/auth/pb"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/models"
)

func AdminLogin(ctx *gin.Context, p pb.AuthServiceClient) {
	var AdminLogin pb.AdminLoginRequest
	err := ctx.BindJSON(&AdminLogin)

	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error parsing the request body", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}

	//lets call the AuthserviceClient for adminLogin
	res, err := p.AdminLogin(ctx, &AdminLogin)

	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error while calling the AdminLogin from api gateway,error connecting the auth service ", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	successRes := models.MakeResponse(http.StatusOK, "admin successfully logged in ", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}
