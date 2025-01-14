package controller

import (
	"e-book/app/service"
	"fmt"
	"net/http"
)

type BookContrller interface {
	CreateBook(w http.ResponseWriter, r *http.Request)
}

type bookContrllerImpl struct {
	bookService service.BookService
}

func NewBookController(bookService service.BookService) BookContrller {
	return &bookContrllerImpl{
		bookService: bookService,
	}
}

func (c *bookContrllerImpl) CreateBook(w http.ResponseWriter, r *http.Request) {
	res, err := c.bookService.CreateBookService(r)

	if err != nil {
		fmt.Println("not sucess", err)
		return
	}
	fmt.Println("response from controller", res)

	// api.Send(w, http.StatusOK, "success")
}
