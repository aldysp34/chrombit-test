package main

import (
	"aldysp34/chrombit-test/handler"
	"aldysp34/chrombit-test/middleware"
	"aldysp34/chrombit-test/usecase"

	"github.com/gin-gonic/gin"
)

type routerOpts struct {
	blogHandler *handler.BlogHandler
	authHandler *handler.AuthHandler
}

func main() {
	blogUsecase := usecase.NewBlogUsecase()
	blogHandler := handler.NewBlogHandler(blogUsecase)
	opts := routerOpts{
		blogHandler: blogHandler,
	}
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middleware.AuthorizeHandler())
	router.Use(middleware.ErrorHandler())

	blogs := router.Group("/api/v1/blogs")
	blogs.GET("", opts.blogHandler.GetBlogs)
	blogs.GET("/:id", opts.blogHandler.GetBlogByID)
	blogs.POST("", opts.blogHandler.CreateBlog)
	blogs.PUT("/:id", opts.blogHandler.EditBlog)
	blogs.DELETE("/:id", opts.blogHandler.DeleteBlog)

	// auth := router.Group("/api/v1/auth")
	// auth.POST("/login", opts.authHandler.Login)
	// auth.POST("/register", opts.authHandler.Register)

	router.Run(":8080")
}
