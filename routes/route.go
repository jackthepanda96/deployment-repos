package routes

import (
	"project/empty_deployment/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterPath(ec *echo.Echo, bc *controllers.BookController) {

	ec.POST("/books", bc.InsertBook)
	ec.GET("/books", bc.GetAllBook)
}
