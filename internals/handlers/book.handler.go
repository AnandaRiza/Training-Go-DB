package handlers

import (
	"bcas/bookstore-go/internals/models"
	"bcas/bookstore-go/internals/repositories"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	*repositories.BookRepo
}

func InitBookHandler(b *repositories.BookRepo) *BookHandler {
	return &BookHandler{b}
}

func (b *BookHandler) GetBooks(ctx *gin.Context) {
	result, err := b.FindAll()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Get Book",
		"data":    result,
	})
}

func (b *BookHandler) CreateBooks(ctx *gin.Context) {
	body := models.BookModel{}
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return

	}

	if err := b.SaveBook(body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := b.FindAll()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success save book",
		"data":    result,
	})
}

// GetBookById retrieves a book by its ID from the database.
func (item *BookHandler) GetBookById(ctx *gin.Context) {
	// Extract the ID from the path variable and convert it to an integer.
	id, _ := strconv.Atoi(ctx.Param("id"))

	// Find the book by its ID.
	result, err := item.FindbyId(id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// If the book is not found, return a "not found" error.
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messages": "book not found",
		})
		return
	}

	// Send the response in JSON format, using gin.H to create a map with string keys and any values.
	ctx.JSON(http.StatusOK, gin.H{
		"messages": "success get book",
		"data":     result,
	})
}

// DeleteBookById deletes a book by its ID from the database.
func (item *BookHandler) DeleteBookById(ctx *gin.Context) {
	// Extract the ID from the path variable and convert it to an integer.
	id, _ := strconv.Atoi(ctx.Param("id"))

	// Find the book by its ID.
	result, err := item.FindbyId(id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// If the book is not found, return a "not found" error.
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messages": "book not found",
		})
		return
	}

	// Delete the book by its ID.
	if err := item.DeletebyId(id); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Send the response in JSON format, using gin.H to create a map with string keys and any values.
	ctx.JSON(http.StatusOK, gin.H{
		"messages": "success delete book",
	})

}

// UpdateBookById updates a book by its ID.
func (item *BookHandler) UpdateBookById(ctx *gin.Context) {
	// Extract the ID from the path variable and convert it to an integer.
	id, _ := strconv.Atoi(ctx.Param("id"))

	// Create a struct to hold the request body.
	body := models.BookModel{}

	// Bind the request body to the struct.
	if err := ctx.ShouldBind(&body); err != nil {
		handleError(ctx, err)
		return
	}

	// Check if any field in the request body is empty.
	if !isValidBookBody(body) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "empty field, At least one field must be provided",
		})
		return
	}

	// Find the book by its ID.
	result, err := item.FindbyId(id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	// If the book is not found, return a "not found" error.
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messages": "book not found",
		})
		return
	}

	// Update the book by its ID.
	if err := item.UpdateById(id, body); err != nil {
		handleError(ctx, err)
		return
	}

	// Send the response in JSON format.
	ctx.JSON(http.StatusOK, gin.H{
		"messages": "success update book",
	})
}

// isValidBookBody checks if the book body contains at least one non-empty field.
func isValidBookBody(body models.BookModel) bool {
	return body.Title != "" || (body.Description != nil && *body.Description != "") || body.Author != ""
}

// handleError handles the error and sends an appropriate response.
func handleError(ctx *gin.Context, err error) {
	log.Println(err.Error())
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}
