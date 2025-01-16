package controller

import (
	"e-book/app/service"
	"e-book/pkg/api"
	"e-book/pkg/e"
	"net/http"
)

type BookController interface {
	CreateBook(w http.ResponseWriter, r *http.Request)
	GetBookById(w http.ResponseWriter, r *http.Request)
	GetallBookDetails(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	DeleteBookById(w http.ResponseWriter, r *http.Request)
}

type bookControllerImpl struct {
	bookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &bookControllerImpl{
		bookService: bookService,
	}
}

func (c *bookControllerImpl) CreateBook(w http.ResponseWriter, r *http.Request) {
	res, err := c.bookService.CreateBookService(r)

	if err != nil {
		apiErr := e.NewAPIError(err, "error while creating book")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, res)
}

func (c *bookControllerImpl) GetBookById(w http.ResponseWriter, r *http.Request) {
	res, err := c.bookService.GetBookById(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "error while getting book detail")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, res)
}

func (c *bookControllerImpl) GetallBookDetails(w http.ResponseWriter, r *http.Request) {
	res, err := c.bookService.GetallBookDetails(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "failed to get All book details")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, res)
}

func (c *bookControllerImpl) UpdateBook(w http.ResponseWriter, r *http.Request) {
	err := c.bookService.UpdateBook(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "failed to update book")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "success")
}

func (c *bookControllerImpl) DeleteBookById(w http.ResponseWriter, r *http.Request) {
	err := c.bookService.DeleteBookById(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "failed to delete book")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "success")
}
