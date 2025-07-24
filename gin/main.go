package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	db := setupDatabase()
	bookapi := r.Group("/api/books")
	bookapi.POST("", func(ctx *gin.Context) {
		book := new(Book)
		if err := ctx.ShouldBindJSON(book); err != nil {
			Faild(ctx, err)
			return
		}

		if err := db.Save(book).Error; err != nil {
			Faild(ctx, err)
		}
		ctx.JSON(http.StatusOK, book)
	})

	bookapi.GET("", func(ctx *gin.Context) {
		var books []Book
		if err := db.Find(&books).Error; err != nil {
			Faild(ctx, err)
			return
		}
		// fmt.Printf("查询结果: %+v\n", books)
		ctx.JSON(http.StatusOK, books)
	})

	bookapi.GET("/:isbn", func(ctx *gin.Context) {
		var book Book
		id := ctx.Param("isbn")
		if err := db.First(&book, id).Error; err != nil {
			Faild(ctx, fmt.Errorf("Book not found"))
			return
		}
		ctx.JSON(http.StatusOK, book)
	})

	bookapi.PUT("/:isbn", func(ctx *gin.Context) {
		id := ctx.Param("isbn")
		req := BookSpec{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			Faild(ctx, err)
			return
		}
		if err := db.Where("isbn=?", id).Model(&Book{}).Updates(req).Error; err != nil {
			Faild(ctx, err)
			return
		}

		var book Book
		if err := db.Where("isbn=?", id).Take(&book).Error; err != nil {
			Faild(ctx, fmt.Errorf("Book not found"))
			return
		}
		ctx.JSON(http.StatusOK, book)
	})

	bookapi.DELETE("/:isbn", func(ctx *gin.Context) {
		id := ctx.Param("isbn")
		if err := db.Where("isbn=?", id).Delete(&Book{}).Error; err != nil {
			Faild(ctx, err)
			return
		}
	})

	err := r.Run("21.6.70.96:8080")
	if err != nil {
		log.Println(err)
	}
}

type Book struct {
	IsBN uint `json:"isbn" gorm:"primaryKey;column:isbn"`
	BookSpec
}

type BookSpec struct {
	Title  string  `json:"title" gorm:"column:title;type:varchar(200)"`
	Author string  `json:"author" gorm:"column:author;type:varchar(200);index"`
	Price  float64 `json:"price" gorm:"column:price"`
	IsSale *bool   `json:"is_sale" gorm:"column:is_sale"`
}

func setupDatabase() *gorm.DB {
	dsn := "root:123456789@tcp(21.6.70.96:3306)/test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai&allowNativePasswords=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败:%v", err))
	}
	db.AutoMigrate(&Book{})
	return db.Debug()
}

func (t *Book) TableName() string {
	return "books"
}

func Faild(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": err.Error()})
}
