package api

import (
	"github.com/xiao-hub-create/book/config"
	"gorm.io/gorm"
)

func NewUserHandler() *UserApiHandler {
	return &UserApiHandler{
		db: config.Get().MySQL.DB(),
	}
}

type UserApiHandler struct {
	db *gorm.DB
}
