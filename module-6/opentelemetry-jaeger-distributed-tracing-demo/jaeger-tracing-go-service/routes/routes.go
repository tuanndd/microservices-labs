package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jaeger-tracing-go-service/controllers"
)

func Routes(router *gin.Engine) {
	router.POST("/employee", controllers.CreateEmployee)
	router.GET("/employee/:id", controllers.GetEmployee)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}