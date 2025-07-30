package controller_test

import (
	"context"
	"testing"

	"github.com/xiao-hub-create/book/controller"
	"github.com/xiao-hub-create/book/models"
)

func TestGetBook(t *testing.T) {
	book := controller.NewBookController()
	ins, err := book.GetBook(context.Background(), &controller.GetBookRequest{
		Isbn: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestCreateBook(t *testing.T) {
	book := controller.NewBookController()
	ins, err := book.CreateBook(context.Background(), &models.BookSpec{
		Author: "kun",
		Price:  10,
		Title:  "啦啦",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)

}
