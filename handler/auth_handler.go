package handler

import (
	"aldysp34/chrombit-test/apperror"
	"aldysp34/chrombit-test/dto"
	"aldysp34/chrombit-test/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(au usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: au}
}

func (ah *AuthHandler) Login(ctx *gin.Context) {
	var loginInfo dto.LoginRequest
	err := ctx.ShouldBindJSON(&loginInfo)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}

	response, err := ah.authUsecase.Login(ctx, loginInfo)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{Data: response, Message: "Login Successfully"})
}
