package service

import (
	"cleanapp/BookStore/intf"
	"cleanapp/BookStore/model"
	"context"
	"fmt"
)

type BookServiceImpl struct {
	BookDatalayer intf.BookDatalayer
}

// This function will create an BookService object which represents the BookService interface
func NewBookServiceImpl(bookDataLayer intf.BookDatalayer) intf.BookService {

	return &BookServiceImpl{
		BookDatalayer: bookDataLayer,
	}

}

func (service *BookServiceImpl) PrintBookTitle(ctx context.Context, book *model.Book) {

}

func (service *BookServiceImpl) TestBookService(ctx context.Context) {

	fmt.Println("**** Inside Book Service ****")

	service.BookDatalayer.TestBookDatalayer(ctx)

}
