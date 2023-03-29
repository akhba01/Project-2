package controllers

import (
	"Project-2/database"
	"Project-2/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBook(ctx *gin.Context) {
	var book []models.Book
	database.GetDB().Find(&book)

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

func GetBookById(ctx *gin.Context) {
	var book models.Book
	db := database.GetDB()

	if err := db.Where("id= ?", ctx.Param("IdBook")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data Tidak Ditemukan",
		})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

type CreateBookInput struct {
	NameBook string `json:"name_book" binding:"required"`
	Author   string `json:"author" binding:"required"`
}

func CreateBook(ctx *gin.Context) {
	var db = database.GetDB()
	var input CreateBookInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create Book
	book := models.Book{NameBook: input.NameBook, Author: input.Author}
	err := db.Create(&book).Error
	if err != nil {
		fmt.Println("Error Creating User data", err)
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

type UpdateBookInput struct {
	NameBook string `json:"name_book"`
	Author   string `json:"author"`
}

func UpdateBook(ctx *gin.Context) {
	book := models.Book{}
	var db = database.GetDB()

	if err := db.Where("id= ?", ctx.Param("IdBook")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	var input UpdateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := db.Model(&book).Where("id= ?", ctx.Param("IdBook")).Updates(models.Book{
		NameBook: input.NameBook,
		Author:   input.Author,
	}).Error
	if err != nil {
		fmt.Println("Error updating book data : ", err)
	}

	ctx.JSON(http.StatusOK, book)

}

func DeleteBook(ctx *gin.Context) {
	var book models.Book
	db := database.GetDB()
	if err := db.Where("id= ?", ctx.Param("IdBook")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return

	}
	err := db.Delete(&book).Error
	if err != nil {
		fmt.Println("Error deleting book data : ", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book Deleted successfully",
	})

}
