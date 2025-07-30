package api

import (
	"github.com/xiao-hub-create/book/config"
	"gorm.io/gorm"
)

func NewCommentHandler() *CommentApiHandler {
	return &CommentApiHandler{
		db: config.Get().MySQL.DB(),
	}
}

type CommentApiHandler struct {
	db *gorm.DB
}
