package dto

import (
	"e-book/pkg/e"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

type BookDeleteRequest struct {
	UserID int64 `json:"userid"`
	BookId int64 `json:"bookid"`
}

func (args *BookDeleteRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return e.NewError(e.ErrResourceNotFound, "convertion error", err)
	}
	args.BookId = int64(intID)
	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		return err
	}

	return nil
}

func (args *BookDeleteRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(args)
	if err != nil {
		return err
	}
	return nil
}
