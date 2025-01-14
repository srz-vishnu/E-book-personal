package dto

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

// struct to hold the data for creating a book
type BookInputRequest struct {
	Title     string `json:"title"`
	AuthorID  int64  `json:"authorid"`
	CreatedBy int64  `json:"createdby"`
	Status    int64  `json:"status"`
	Content   string `json:"content"`
}

type BookOutputResponse struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

func (args *BookInputRequest) Parse(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&args)
	if err != nil {
		return err
	}
	return nil
}

func (args *BookInputRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(args)
	if err != nil {
		return err
	}
	return nil
}
