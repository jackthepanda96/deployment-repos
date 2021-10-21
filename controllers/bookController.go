package controllers

import (
	"fmt"
	"net/http"

	"project/empty_deployment/models"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	bookModel models.BookModel
}

func NewBookController(bm models.BookModel) *BookController {
	return &BookController{bm}
}

func (b *BookController) InsertBook(ec echo.Context) error {
	book := models.Book{}
	ec.Bind(&book)
	fmt.Println(book)
	book, err := b.bookModel.Insert(book)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, "invalid book data")
	}
	return ec.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success add book",
		"data":    book,
	})
}

func (b *BookController) GetAllBook(ec echo.Context) error {
	books, err := b.bookModel.GetAll()

	if err != nil {
		return ec.JSON(http.StatusInternalServerError, "cannot get books")
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all book",
		"data":    books,
	})
}
