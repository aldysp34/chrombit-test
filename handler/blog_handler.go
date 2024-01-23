package handler

import (
	"aldysp34/chrombit-test/apperror"
	"aldysp34/chrombit-test/dto"
	"aldysp34/chrombit-test/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	blogUsecase usecase.BlogUsecase
}

func NewBlogHandler(blogUsecase usecase.BlogUsecase) *BlogHandler {
	return &BlogHandler{
		blogUsecase: blogUsecase,
	}
}

const (
	SuccessMessage = "successfully"
)

func (bh *BlogHandler) GetBlogs(ctx *gin.Context) {
	resp := bh.blogUsecase.GetBlogs(ctx)
	ctx.JSON(http.StatusOK, dto.Response{Data: resp, Message: SuccessMessage})
}

func (bh *BlogHandler) GetBlogByID(ctx *gin.Context) {
	blogID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(apperror.ErrAtoiString)
		return
	}

	resp, err := bh.blogUsecase.GetBlogByID(ctx, blogID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{Data: resp, Message: SuccessMessage})
}

func (bh *BlogHandler) CreateBlog(ctx *gin.Context) {
	var data dto.BlogRequest

	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}

	blogCreated, err := bh.blogUsecase.CreateBlog(ctx, data)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{Data: blogCreated, Message: SuccessMessage})
}

func (bh *BlogHandler) EditBlog(ctx *gin.Context) {
	var data dto.BlogRequest

	blogID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(apperror.ErrAtoiString)
		return
	}
	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	data.ID = blogID

	response, err := bh.blogUsecase.EditBlog(ctx, data)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{Data: response, Message: SuccessMessage})
}

func (bh *BlogHandler) DeleteBlog(ctx *gin.Context) {
	blogID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(apperror.ErrAtoiString)
		return
	}

	err = bh.blogUsecase.DeleteBlog(ctx, blogID)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
