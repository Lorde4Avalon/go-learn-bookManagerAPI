package routers

import (
	"github.com/Rioa-Avalon/go-demo/controllers"
	"github.com/gin-gonic/gin"
)

func RouterRegist() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/book")
	{
		v1.POST("/new", controllers.CreateBook)
		v1.GET("/get/all", controllers.GetAllBooks)
		v1.GET("/get/:id", controllers.GetBookById)
		v1.DELETE("/delete/:id", controllers.DeleteBookById)
	}
	return router
}
