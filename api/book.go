package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xiao-hub-create/book/config"
	"github.com/xiao-hub-create/book/controller"
	"github.com/xiao-hub-create/book/models"
	"github.com/xiao-hub-create/book/response"
	"gorm.io/gorm"
)

// 构造函数，初始化结构体
func NewBookHandler() *BookApiHandler {
	return &BookApiHandler{
		db:  *config.Get().MySQL.DB(),
		svc: *controller.NewBookController(),
	}
}

// 面向对象
type BookApiHandler struct {
	db  gorm.DB
	svc controller.BookController
}

// 提供注册功能，提供一个Group
func (h *BookApiHandler) Registry(r *gin.Engine) {
	book := r.Group("/api/books")
	book.POST("", h.CreateBook)
	book.GET("", h.ListBook)
	book.GET("/:isbn", h.GetBook)
	book.PUT("/:isbn", h.UpdateBook)
	book.DELETE("/:isbn", h.DeleteBook)
}

func (h *BookApiHandler) CreateBook(ctx *gin.Context) {
	req := new(models.BookSpec)
	if err := ctx.ShouldBindJSON(req); err != nil {
		response.Failed(ctx, err)
		return
	}

	ins, err := h.svc.CreateBook(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)
}

func (h *BookApiHandler) ListBook(ctx *gin.Context) {

	var books []models.Book
	if err := h.db.Find(&books).Error; err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, books)

}

func (h *BookApiHandler) GetBook(ctx *gin.Context) {
	strid := ctx.Param("isbn")
	id, err := strconv.ParseInt(strid, 10, 64)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	ins, err := h.svc.GetBook(ctx.Request.Context(), &controller.GetBookRequest{
		Isbn: id,
	})
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)
}

func (h *BookApiHandler) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("isbn")
	req := models.BookSpec{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Failed(ctx, err)
		return
	}
	if err := h.db.Where("isbn=?", id).Model(&models.Book{}).Updates(req).Error; err != nil {
		response.Failed(ctx, err)
		return
	}

	var book models.Book
	if err := h.db.Where("isbn=?", id).Take(&book).Error; err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, book)
}

func (h *BookApiHandler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("isbn")
	if err := h.db.Where("isbn=?", id).Delete(models.Book{}).Error; err != nil {
		response.Failed(ctx, err)
		return
	}
}

// bookapi := r.Group("/api/books")
// bookapi.POST("", func(ctx *gin.Context) {
// 	book := new(models.Book)
// 	if err := ctx.ShouldBindJSON(book); err != nil {
// 		Faild(ctx, err)
// 		return
// 	}

// 	if err := db.Save(book).Error; err != nil {
// 		Faild(ctx, err)
// 	}
// 	ctx.JSON(http.StatusOK, book)
// })

// bookapi.GET("", func(ctx *gin.Context) {
// 	var books []models.Book
// 	if err := db.Find(&books).Error; err != nil {
// 		Faild(ctx, err)
// 		return
// 	}
// 	// fmt.Printf("查询结果: %+v\n", books)
// 	ctx.JSON(http.StatusOK, books)
// })

// bookapi.GET("/:isbn", func(ctx *gin.Context) {
// 	var book models.Book
// 	id := ctx.Param("isbn")
// 	if err := db.First(&book, id).Error; err != nil {
// 		Faild(ctx, fmt.Errorf("Book not found"))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, book)
// })

// bookapi.PUT("/:isbn", func(ctx *gin.Context) {

// })

// bookapi.DELETE("/:isbn", func(ctx *gin.Context) {
// 	id := ctx.Param("isbn")
// 	if err := db.Where("isbn=?", id).Delete(&models.Book{}).Error; err != nil {
// 		Faild(ctx, err)
// 		return
// 	}
// })
