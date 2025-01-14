package dto

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type BookDetailsById struct {
	BookID int64 `json:"bookid"`
}

type BookDetailsByIdResponse struct {
	Title string `json:"title"`
}

func (args *BookDetailsById) Parse(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&args)
	if err != nil {
		return err
	}
	return nil
}

func (args *BookDetailsById) Validate() error {
	validate := validator.New()
	err := validate.Struct(args)
	if err != nil {
		return err
	}
	return nil
}
