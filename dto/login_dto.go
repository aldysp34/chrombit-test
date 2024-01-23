package dto

type LoginRequest struct {
	Email    string `binding:"required,email" json:"email" validate:"required,email"`
	Password string `binding:"required" json:"password" validate:"required,min=8,max=35"`
}
type LoginResponse struct {
	Token string `json:"token"`
}
