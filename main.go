package main

import (
	"cleanapp/BookStore/controller"
	"cleanapp/BookStore/database"
	"cleanapp/BookStore/service"

	"github.com/labstack/echo"
)

func main() {

	echoContext := echo.New()

	datalayer := database.NewBookDatalayerImpl()

	service := service.NewBookServiceImpl(datalayer)

	controller.NewBookController(echoContext, service)

	echoContext.Logger.Info(echoContext.Start(":8090"))

}
