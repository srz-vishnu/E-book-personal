package dto

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

type GetAuthorDetailRequest struct {
	AuthorId int64 `json:"authorid"`
}

type GetAuthorDetailResponse struct {
	AuthorName string `json:"authorname"`
	AuthorId   int64  `json:"authorid"`
}

func (args *GetAuthorDetailRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	args.AuthorId = int64(intID)
	return nil
}

func (args *GetAuthorDetailRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(args)
	if err != nil {
		return err
	}
	return nil
}
