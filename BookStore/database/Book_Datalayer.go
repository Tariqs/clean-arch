package database

import (
	"cleanapp/BookStore/intf"
	"context"
	"fmt"
)

type BookDatalayerImpl struct {
}

func NewBookDatalayerImpl() intf.BookDatalayer {

	return &BookDatalayerImpl{}

}

func (datalayer *BookDatalayerImpl) GetBookByID(ctx context.Context, bookID int16) {

}

func (datalayer *BookDatalayerImpl) TestBookDatalayer(ctx context.Context) {

	fmt.Println("**** Inside Book Datalayer ****")
}
