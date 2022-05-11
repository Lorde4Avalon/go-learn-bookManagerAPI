package controllers

import (
	"fmt"
	"net/http"

	"github.com/Rioa-Avalon/go-demo/models"
	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.GetAllBooks())
}

func CreateBook(c *gin.Context) {
	var newbook models.Book
	c.Bind(&newbook)
	fmt.Println(newbook)
	err := models.CreateBook(newbook)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "isnb is unique!!!",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "new book create success!",
		})
	}
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	book, _ := models.GetBookById(id)
	
	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "book not found",
		})
	} else {
		c.JSON(http.StatusOK, book)
	}

}

func DeleteBookById(c *gin.Context) {
	id := c.Param("id")
	book, _ := models.DeleteBookById(id)

	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "book not found",
		})
	} else {
		c.IndentedJSON(http.StatusOK, book)
	}
}
