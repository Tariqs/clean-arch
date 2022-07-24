package controller

import (
	"cleanapp/BookStore/intf"
	"fmt"

	"github.com/labstack/echo"
)

type BookController struct {
	bookService intf.BookService
}

func NewBookController(echo *echo.Echo, bookServiceObject intf.BookService) {

	bookControllerObject := &BookController{
		bookService: bookServiceObject,
	}

	echo.GET("/printAuthor", bookControllerObject.PrintAuthor)
	echo.GET("/test", bookControllerObject.Test)
}

func (controllerObj *BookController) PrintAuthor(ec echo.Context) error {

	return nil
}

func (controllerObj *BookController) Test(ec echo.Context) error {

	fmt.Println("**** Inside Book Controller ****")

	requestContext := ec.Request().Context()
	controllerObj.bookService.TestBookService(requestContext)

	return nil
}
