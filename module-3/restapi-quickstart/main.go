package main

import (
	"demo/controllers"
	"demo/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	// })

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
