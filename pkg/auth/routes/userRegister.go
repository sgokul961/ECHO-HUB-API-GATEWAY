package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/auth/pb"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/helper"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/models"
)

// @Summary		Register a new user
// @Description	Register a new user with the provided information
// @Tags			Authentication
// @Accept			multipart/form-data
// @Produce		json
// @Param			email			formData	string			true	"User's email address"
// @Param			password		formData	string			true	"User's password"
// @Param			phonenum		formData	string			true	"User's phone number"
// @Param			username		formData	string			true	"User's username"
// @Param			bio				formData	string			false	"User's bio"
// @Param			gender			formData	string			false	"User's gender"
// @Param			profile_picture	formData	file			true	"User's profile picture"
// @Success		200				{object}	models.Response	"Success"
// @Failure		400				{object}	models.Response	"Bad Request"
// @Failure		502				{object}	models.Response	"Bad Gateway"
// @Router			/auth/register [post]
func Register(ctx *gin.Context, c pb.AuthServiceClient) {

	err := ctx.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var req pb.RegisterRequest

	req.Email = ctx.Request.FormValue("email")
	req.Password = ctx.Request.FormValue("password")
	req.Phonenum = ctx.Request.FormValue("phonenum")
	req.Username = ctx.Request.FormValue("username")
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
	//call the register service to handle the register function
	res, err := c.Register(ctx, &req)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error in connecting with service", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	succRes := models.MakeResponse(http.StatusOK, "successfully registered user data", res, nil)

	ctx.JSON(http.StatusOK, succRes)

}
