package models

import (
	"errors"
	"time"

	"github.com/Rioa-Avalon/go-demo/config"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

var db *gorm.DB

type Book struct {
	ID        uint `gorm:"primarykey"`
    	CreatedAt time.Time
    	UpdatedAt time.Time
    	Is_Del soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;default:0"`
	Title  string `gorm:""json:"title"`
	Author string `json:"author"`
	Isbn   string `gorm:"uniqueIndex:udx_name;type:varchar(30)"json:"isbn"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func CreateBook(b Book) error {
	result := db.Create(&b)
	return result.Error
}

func GetAllBooks() []Book {
	var Books []Book
	db.Where("is_Del=?", "0").Find(&Books)
	return Books
}

func GetBookById(Id string) (*Book, error) {
	var book Book
	result := db.Where("Id=? AND Is_Del=?", Id, "0").Find(&book)

	if result.RowsAffected == 0 {
		return nil, errors.New("book not found")
	}
	return &book, result.Error
}

func DeleteBookById(Id string) (*Book, error) {
	var book Book
	result := db.Model(&book).Where("Id=?", Id).Update("Is_Del", 1)
	if result.RowsAffected == 0 {
		return nil, errors.New("book not found")
	}
	return &book, result.Error
}
