package middleware

import (
	"aldysp34/chrombit-test/apperror"
	"aldysp34/chrombit-test/dto"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if strings.HasPrefix(ctx.Request.URL.Path, "/api/v1/blogs") {
			ctx.Next()
			return
		}
		if strings.HasPrefix(ctx.Request.URL.Path, "/api/v1/auth") {
			ctx.Next()
			return
		}

		var response dto.Response
		header := ctx.Request.Header["Authorization"]
		if len(header) == 0 {
			response.Message = apperror.ErrInvalidToken.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		splittedHeader := strings.Split(header[0], " ")
		if len(splittedHeader) != 2 {
			response.Message = apperror.ErrUnauthorized.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claims := &dto.JWTClaims{}

		token, err := jwt.ParseWithClaims(splittedHeader[1], claims, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, apperror.ErrInvalidCredentials
			}

			return []byte("secret"), nil
		})
		if err != nil {
			response.Message = err.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		_, ok := token.Claims.(*dto.JWTClaims)
		if !ok || !token.Valid {
			response.Message = apperror.ErrUnauthorized.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		newContext := context.WithValue(ctx.Request.Context(), "user_id", claims.UserID)
		ctx.Request = ctx.Request.WithContext(newContext)
		ctx.Next()
	}

}
