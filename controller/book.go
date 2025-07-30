package controller

import (
	"context"

	"github.com/xiao-hub-create/book/config"
	"github.com/xiao-hub-create/book/models"
	"gorm.io/gorm"
)

func NewBookController() *BookController {
	return &BookController{
		db: config.Get().MySQL.DB(),
	}
}

type BookController struct {
	db *gorm.DB
}

type GetBookRequest struct {
	Isbn int64
}

// 这里是纯业务处理，和http没关系
func (c *BookController) GetBook(ctx context.Context, req *GetBookRequest) (*models.Book, error) {

	ins := &models.Book{}

	if err := c.db.WithContext(ctx).Where("isbn=?", req.Isbn).Take(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (c *BookController) CreateBook(ctx context.Context, req *models.BookSpec) (*models.Book, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	ins := &models.Book{
		BookSpec: *req,
	}

	if err := c.db.Save(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}
