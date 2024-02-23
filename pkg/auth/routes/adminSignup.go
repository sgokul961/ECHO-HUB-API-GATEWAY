package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/auth/pb"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/helper"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/models"
)

func AdminSignup(ctx *gin.Context, p pb.AuthServiceClient) {

	err := ctx.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var req pb.AdminSignupRequest
	req.Email = ctx.Request.FormValue("email")
	req.Password = ctx.Request.FormValue("password")
	req.Username = ctx.Request.FormValue("username")
	req.Phonenum = ctx.Request.FormValue("phonenum")
	req.Bio = ctx.Request.FormValue("bio")
	req.Gender = ctx.Request.FormValue("gender")
	file, fileHeader, err := ctx.Request.FormFile("profile_picture")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer file.Close()
	if fileHeader == nil {
		errMsg := "profile_picture is required"
		ctx.AbortWithError(http.StatusBadRequest, errors.New(errMsg))
		return
	}
	picture, err := helper.AddImageToS3(fileHeader)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	req.ProfilePicture = picture

	res, err := p.AdminSignup(ctx, &req)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error in connecting with service", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	succRes := models.MakeResponse(http.StatusOK, "successfully registered admin data", res, nil)

	ctx.JSON(http.StatusOK, succRes)

}
