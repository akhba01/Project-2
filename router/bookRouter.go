package router

import (
	"Project-2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.GetAllBook)

	router.GET("/books/:IdBook", controllers.GetBookById)

	router.POST("/books", controllers.CreateBook)

	router.PUT("/books/:IdBook", controllers.UpdateBook)

	router.DELETE("books/:IdBook", controllers.DeleteBook)

	return router
}
