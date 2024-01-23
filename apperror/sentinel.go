package apperror

import "net/http"

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func (c *CustomError) Error() string {
	return c.Message
}

func (c *CustomError) ConvertToErrorResponse() ErrorResponse {
	return ErrorResponse{
		Message: c.Message,
	}
}

var (
	ErrAtoiString         = NewCustomError(http.StatusBadRequest, "can't convert string to integer")
	ErrInvalidBody        = NewCustomError(http.StatusBadRequest, "invalid body")
	ErrBlogNotFound       = NewCustomError(http.StatusNotFound, "blog id not found")
	ErrDuplicateBlog      = NewCustomError(http.StatusBadRequest, "blog id is duplicate")
	ErrInvalidToken       = NewCustomError(http.StatusUnauthorized, "bearer token is missing")
	ErrUnauthorized       = NewCustomError(http.StatusUnauthorized, "unauthorized")
	ErrInvalidCredentials = NewCustomError(http.StatusUnauthorized, "invalid credentials")
)
