package dto

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type BookDeleteRequest struct {
	UserID int64 `json:"userid"`
	BookId int64 `json:"bookid"`
}

func (args *BookDeleteRequest) Parse(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&args)
	if err != nil {
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
