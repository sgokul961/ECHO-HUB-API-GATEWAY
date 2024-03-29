package routes

import (
	"errors"
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
func ResetPassword(ctx *gin.Context, p pb.AuthServiceClient) {
	var Reset pb.ResetPasswordRequest
	id, ok := ctx.Get("userId")
	if !ok {
		errRes := models.MakeResponse(http.StatusUnauthorized, "error parsing id", nil, errors.New("error in id retrieval").Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}

	err := ctx.BindJSON(&Reset)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error parsing request body,email ", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	fmt.Println("rest", Reset.Email, Reset.Password)

	Reset.Id = id.(int64)

	res, err := p.ResetPassword(ctx, &Reset)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error connecting auth service", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	fmt.Println("res", res.Error)
	successRes := models.MakeResponse(http.StatusOK, " new password successfully updated", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}
func ForgotPassWord(ctx *gin.Context, p pb.AuthServiceClient) {
	var forgot pb.ForgotPasswordRequest

	err := ctx.BindJSON(&forgot)

	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error parsing request body,email ", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	res, err := p.ForgotPassWord(ctx, &forgot)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error connecting auth service", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	successRes := models.MakeResponse(http.StatusOK, " forgot password successfully updated", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}
func CheckUserBlocked(ctx *gin.Context, p pb.AuthServiceClient) {
	var checkblock pb.CheckUserBlockedRequest

	err := ctx.BindJSON(&checkblock)

	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error parsing request body,email ", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	res, err := p.CheckUserBlocked(ctx, &checkblock)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error connecting auth service", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	if res.Status {
		fmt.Println("blocked")
	} else {
		fmt.Println("not blocked")
	}

	successRes := models.MakeResponse(http.StatusOK, "User blocked status checked successfully", res.Status, nil)
	ctx.JSON(http.StatusOK, successRes)

}
func FetchShortDetails(ctx *gin.Context, p pb.AuthServiceClient) {

	var fetch pb.FetchShortDetailsRequest

	err := ctx.BindJSON(&fetch)

	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error parsing request body,id ", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}

	res, err := p.FetchShortDetails(ctx, &fetch)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error connecting auth service", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	fmt.Println("id:", res.Id, "image:", res.Image, "name:", res.Image)
	// response := pb.FetchShortDetailsResponse{
	// 	Id:    res.Id,
	// 	Name:  res.Name,
	// 	Image: res.Image,
	// }
	successRes := models.MakeResponse(http.StatusOK, "Successfyllly got all details", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}
