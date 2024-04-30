package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/auth/pb"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/models"
)

// AdminLogin godoc
// @Summary Admin login
// @Description Authenticate admin user
// @Accept json
// @Produce json
// @Param request body pb.AdminLoginRequest true "Admin login request body"
// @Success 200 {object} string "Admin successfully logged in"
// @Failure 502 {object} string "Error parsing request body or connecting to authentication service"
// @Router /admin/adminlogin [post]
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

// BlockUser godoc
// @Summary Block user
// @Description Block a user account
// @Accept json
// @Produce json
// @Param request body pb.BlockUserRequest true "Block user request body"
// @Success 200 string SuccessResponse "Admin successfully blocked user"
// @Failure 502 string ErrorResponse "Error parsing request body or connecting to authentication service"
// @Router /admin/block [post]
func BlockUser(ctx *gin.Context, p pb.AuthServiceClient) {
	var blockUser pb.BlockUserRequest

	err := ctx.BindJSON(&blockUser)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error parsing the request body", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	res, err := p.BlockUser(ctx, &blockUser)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error while calling the Blockuser from api gateway,error connecting the auth service ", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	successRes := models.MakeResponse(http.StatusOK, "admin successfully blocked user  ", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}

// UnblockUser godoc
// @Summary Unblock user
// @Description Unblock a user account
// @Accept json
// @Produce json
// @Param request body pb.UnblockUserRequest true "Unblock user request body"
// @Success 200 string SuccessResponse "Admin successfully unblocked user"
// @Failure 502 string ErrorResponse "Error parsing request body or connecting to authentication service"
// @Router /admin/unblock [post]
func UnblockUser(ctx *gin.Context, p pb.AuthServiceClient) {

	var unBlockUser pb.UnblockUserRequest

	err := ctx.BindJSON(&unBlockUser)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error parsing the request body", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}
	res, err := p.UnblockUser(ctx, &unBlockUser)
	if err != nil {
		errRes := models.MakeResponse(http.StatusBadGateway, "error while calling the Blockuser from api gateway,error connecting the auth service ", nil, err.Error())
		ctx.JSON(http.StatusBadGateway, errRes)
		return
	}

	successRes := models.MakeResponse(http.StatusOK, "admin successfully unblocked user  ", res, nil)
	ctx.JSON(http.StatusOK, successRes)

}
