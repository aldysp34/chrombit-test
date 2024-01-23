package dto

type BlogRequest struct {
	ID    int    `json:"id"`
	Title string `json:"title" binding:"required" validate:"required"`
	Body  string `json:"body" binding:"required" validate:"required"`
	Slug  string `json:"slug" binding:"required" validate:"required"`
}

type BlogResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Slug  string `json:"slug"`
}
