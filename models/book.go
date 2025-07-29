package models

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

func (t *Book) TableName() string {
	return "books"
}
