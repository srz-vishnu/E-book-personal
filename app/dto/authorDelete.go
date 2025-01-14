package dto

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type AuthorDeleteRequest struct {
	UserID   int64 `json:"userid" validate:"required"`
	AuthorId int64 `json:"authorid" validate:"required"`
}

func (args *AuthorDeleteRequest) Parse(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&args)
	if err != nil {
		return err
	}
	return nil
}

func (args *AuthorDeleteRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(args)
	if err != nil {
		return err
	}
	return nil
}
