package dto

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

type BookDetailsById struct {
	BookID int `json:"id"`
}

type BookDetailsByIdResponse struct {
	Title string `json:"title"`
}

func (args *BookDetailsById) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	args.BookID = intID
	return nil
}

// func (u *UserRequest) Parse(r *http.Request) error {
// 	strID := chi.URLParam(r, "id")
// 	intID, err := strconv.Atoi(strID)
// 	if err != nil {
// 		return err
// 	}
// 	u.ID = intID
// 	return nil
// }

func (args *BookDetailsById) Validate() error {
	validate := validator.New()
	err := validate.Struct(args)
	if err != nil {
		return err
	}
	return nil
}
