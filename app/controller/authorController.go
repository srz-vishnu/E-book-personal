package controller

import (
	"e-book/app/service"
	"net/http"
)

type AuthorController interface {
	CreateAuthor(w http.ResponseWriter, r *http.Request)
}

type authorControllerImpl struct {
	authorService service.AuthorService
}

func NewAuthorController(authorService service.AuthorService) AuthorController {
	return &authorControllerImpl{
		authorService: authorService,
	}
}
