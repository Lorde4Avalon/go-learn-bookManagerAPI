package controllers

import (
	"fmt"
	"net/http"

	"github.com/Rioa-Avalon/go-demo/model"
	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.GetAllBooks())
}

func CreateBook(c *gin.Context) {
	var newbook models.Book
	c.Bind(&newbook)
	fmt.Println(newbook)
	newbook.CreateBook()
	c.JSON(200, gin.H{
		"message": "new book create success!",
	})
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	book, _ := models.GetBookById(id)
	c.IndentedJSON(http.StatusOK, book)
}

func DeleteBookById(c *gin.Context) {
	id := c.Param("id")
	book, _ := models.DeleteBookById(id)
	c.IndentedJSON(http.StatusOK, book)
}
