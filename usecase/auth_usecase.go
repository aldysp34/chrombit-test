package usecase

import (
	"aldysp34/chrombit-test/apperror"
	"aldysp34/chrombit-test/dto"
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Login(context.Context, dto.LoginRequest) (*dto.LoginResponse, error)
}

type authUsecase struct{}

type User struct {
	email    string
	password string
}

var dummyUser User

func NewAuthUsecase() AuthUsecase {
	dummyUser.email = "aldysp33@gmail.com"
	// password = aldysp33
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("aldysp33"), 10)

	dummyUser.password = string(hashedPassword)
	return &authUsecase{}
}

func (au *authUsecase) Login(ctx context.Context, info dto.LoginRequest) (*dto.LoginResponse, error) {
	var response dto.LoginResponse
	if !strings.EqualFold(info.Email, dummyUser.email) {
		return nil, apperror.ErrInvalidCredentials
	}
	err := bcrypt.CompareHashAndPassword([]byte(dummyUser.password), []byte(info.Password))
	if err != nil {
		return nil, apperror.ErrInvalidCredentials
	}

	token, _ := dto.GenerateToken(dto.JWTClaims{
		UserID: 1,
	})
	response.Token = token
	return &response, nil
}

func (au *authUsecase) comparePassword(hashedPwd, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	password := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, password)
	return err == nil
}
