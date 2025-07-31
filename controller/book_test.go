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

func TestListBook(t *testing.T) {
	book := controller.NewBookController()
	ins, err := book.ListBook(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestUpdate(t *testing.T) {
	book := controller.NewBookController()
	ins, err := book.UpdateBook(context.Background(), controller.UpdateBookRequest{
		Isbn: 1,
		BookSpec: models.BookSpec{
			Author: "The King Of Northeast",
			Price:  250,
			Title:  "我爱熊熊",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDelete(t *testing.T) {
	book := controller.NewBookController()
	ins, err := book.DeleteBook(context.Background(), controller.DeleteBookRequest{
		Isbn: 6,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
