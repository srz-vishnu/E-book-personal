package dto

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type CreateAuthorRequest struct {
	AuthorName string `json:"authorname" validate:"required"`
	UserId     int64  `json:"userid" validate:"required"`
}

type CreateAuthorResponse struct {
	AuthorName string `json:"authorname"`
	AuthorId   int64  `json:"authorid"`
}

func (args *CreateAuthorRequest) Parse(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&args)
	if err != nil {
		return err
	}
	return nil
}

func (args *CreateAuthorRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(args)
	if err != nil {
		return err
	}
	return nil
}
