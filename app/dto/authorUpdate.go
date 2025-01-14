package dto

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type AuthorupdateRequest struct {
	AuthorName string `json:"authorname" validate:"required"`
	AuthorId   int64  `json:"authorid" validate:"required"`
	UserID     int64  `json:"userid" validate:"required"`
}

func (args *AuthorupdateRequest) Parse(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&args)
	if err != nil {
		return err
	}
	return nil
}

func (args *AuthorupdateRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(args)
	if err != nil {
		return err
	}
	return nil
}
