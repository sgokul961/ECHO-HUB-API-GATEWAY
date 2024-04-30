package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/auth/pb"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/helper"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/models"
)

// AdminSignup godoc
// @Summary Admin signup
// @Description Register a new admin account
// @Accept multipart/form-data
// @Produce json
// @Param email formData string true "Admin email"
// @Param password formData string true "Admin password"
// @Param username formData string true "Admin username"
// @Param phonenum formData string true "Admin phone number"
// @Param bio formData string false "Admin bio"
// @Param gender formData string false "Admin gender"
// @Param profile_picture formData file true "Admin profile picture"
// @Success 200 string SuccessResponse "Successfully registered admin data"
// @Failure 400 string ErrorResponse "Bad request, error parsing form or missing required fields"
// @Failure 502 string ErrorResponse "Error connecting to authentication service"
// @Router /admin/adminsignup [post]
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
