package dto

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

// bookUpdateRequest is a struct to hold the data for updating a book
type BookUpdateRequest struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  int64  `json:"status"`
	UserID  int64  `json:"userid"`
}

func (args *BookUpdateRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	args.ID = int64(intID)

	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		return err
	}

	return nil
}

func (args *BookUpdateRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(args)
	if err != nil {
		return err
	}
	return nil
}
