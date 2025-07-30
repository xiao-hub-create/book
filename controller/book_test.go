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
		Isbn: 3,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestCreateBook(t *testing.T) {
	book := controller.NewBookController()
	ins, err := book.CreateBook(context.Background(), &models.BookSpec{
		Author: "zeng",
		Price:  20,
		Title:  "牛牛",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)

}
