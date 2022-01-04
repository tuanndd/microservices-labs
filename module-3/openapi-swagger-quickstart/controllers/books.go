package controllers

import (
	"demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// FindBooks godoc
// @Summary List books
// @Description get books
// @Tags books
// @Accept  json
// @Produce json
// @Success 200 {array} models.Book
// @Failure 400 {object} map[string]interface{}
// @Router /books [get]
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// FindBooks godoc
// @Summary Show a book
// @Description get book by ID
// @Tags books
// @Accept  json
// @Produce json
// @Param        id       path      int                  true  "Book ID"
// @Success 200 {object} models.Book
// @Failure 400 {object} map[string]interface{}
// @Router /books/{id} [get]
func FindBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreateBook godoc
// @Summary      Add a book
// @Description  add by json book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        account  body      controllers.CreateBookInput  true  "Create book"
// @Success      200      {object}  models.Book
// @Failure 400 {object} map[string]interface{}
// @Router       /books [post]
func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook godoc
// @Summary      Update a book
// @Description  update by json book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id       path      int                  true  "Book ID"
// @Param        account  body      controllers.UpdateBookInput  true  "Update book"
// @Success      200      {object}  models.Book
// @Failure 400 {object} map[string]interface{}
// @Router       /books/{id} [patch]
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(models.Book{Title: input.Title, Author: input.Author})
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook godoc
// @Summary      Delete a book
// @Description  update by book ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id       path      int                  true  "Book ID"
// @Success      200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router       /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
