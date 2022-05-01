package models

import (
	"errors"

	"github.com/Rioa-Avalon/go-demo/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title  string `gorm:""json:"title"`
	Author string `json:"author"`
	Isbn   string `json:"isbn"`
}

func Init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b Book) CreateBook() {
	db.Create(&b)
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id string) (Book, error) {
	var book Book
	result := db.Where("Id=?", Id).Find(&book)
	return book, result.Error
}

func DeleteBookById(Id string) (Book, error) {
	var book Book
	db.Where("Id=?", Id).Delete(&book)
	return book, errors.New("book not found")
}
