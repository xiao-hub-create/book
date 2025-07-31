package controller

import (
	"context"
	"errors"
	"fmt"

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

type UpdateBookRequest struct {
	Isbn int64
	models.BookSpec
}

type DeleteBookRequest struct {
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

func (c *BookController) ListBook(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	if err := c.db.WithContext(ctx).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (c *BookController) UpdateBook(ctx context.Context, req UpdateBookRequest) (*models.Book, error) {
	var book models.Book
	if err := c.db.WithContext(ctx).Where("isbn=?", req.Isbn).First(&book).Error; err != nil {
		return nil, err
	}

	if err := c.db.WithContext(ctx).Where(req.Isbn).Updates(req.BookSpec).Error; err != nil {
		return nil, err
	}

	if err := c.db.Where("isbn=?", req).Take(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (c *BookController) DeleteBook(ctx context.Context, req DeleteBookRequest) (string, error) {
	var book models.Book
	if err := c.db.WithContext(ctx).Where("isbn=?", req.Isbn).First(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("书籍不存在")
		}
		return "", err
	}
	if err := c.db.WithContext(ctx).Where("isbn=?", req.Isbn).Delete(models.Book{}).Error; err != nil {
		return "", err
	}
	return "删除成功", nil
}
