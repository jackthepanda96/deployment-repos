package main

import (
	"fmt"
	"project/empty_deployment/configs"
	"project/empty_deployment/controllers"
	"project/empty_deployment/middlewares"
	"project/empty_deployment/models"
	"project/empty_deployment/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig()

	db := configs.MysqlConnect(config)

	bookModel := models.NewBookModel(db)

	bookController := controllers.NewBookController(bookModel)

	ec := echo.New()

	routes.RegisterPath(ec, bookController)
	middlewares.GlobalMiddlewares(ec)

	runServer := fmt.Sprint(":", config.Port)

	if err := ec.Start(runServer); err != nil {
		log.Info("shutting down the server", err)
	}
}
