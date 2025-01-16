package controller

import (
	"e-book/app/service"
	"e-book/pkg/api"
	"e-book/pkg/e"
	"net/http"
)

type AuthorController interface {
	CreateAuthor(w http.ResponseWriter, r *http.Request)
	GetAuthorById(w http.ResponseWriter, r *http.Request)
	Updateauthor(w http.ResponseWriter, r *http.Request)
	GetallAuthorInfo(w http.ResponseWriter, r *http.Request)
	DeleteAuthorById(w http.ResponseWriter, r *http.Request)
}

type authorControllerImpl struct {
	authorService service.AuthorService
}

func NewAuthorController(authorService service.AuthorService) AuthorController {
	return &authorControllerImpl{
		authorService: authorService,
	}
}

func (c *authorControllerImpl) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	result, err := c.authorService.CreateAuthor(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "author creation failed")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, result)
}

func (c *authorControllerImpl) GetAuthorById(w http.ResponseWriter, r *http.Request) {
	resp, err := c.authorService.GetAuthorById(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "error while getting author details of the given id")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, resp)
}

func (c *authorControllerImpl) Updateauthor(w http.ResponseWriter, r *http.Request) {
	err := c.authorService.UpdateAuthor(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "error while getting author details of the given id")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "success")
}

func (c *authorControllerImpl) GetallAuthorInfo(w http.ResponseWriter, r *http.Request) {
	resp, err := c.authorService.GetallAuthorDetails(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "error while getting author details")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, resp)
}

func (c *authorControllerImpl) DeleteAuthorById(w http.ResponseWriter, r *http.Request) {
	err := c.authorService.DeleteAuthorById(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "error while deleting the author")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "success")
}
