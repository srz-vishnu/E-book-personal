package dto

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

// bookUpdateRequest is a struct to hold the data for updating a book
type BookUpdateRequest struct {
	BookID  int64  `json:"bookid"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  int64  `json:"status"`
	UserID  int64  `json:"userid"`
}

func (args *BookUpdateRequest) Parse(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&args)
	if err != nil {
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
