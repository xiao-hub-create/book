package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-hub-create/book/config"
	"gorm.io/gorm"
)

func NewCommentController() *CommentController {
	return &CommentController{
		db:   config.Get().MySQL.DB(),
		book: NewBookController(),
	}
}

type CommentController struct {
	db   *gorm.DB
	book *BookController
}

func (h *CommentController) AddComment(ctx gin.Context) {

}
