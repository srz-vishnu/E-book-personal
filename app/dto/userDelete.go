package dto

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type DeleteUserRequest struct {
	UserId int64 `json:"userid" validate:"required"`
}

func (args *DeleteUserRequest) Parse(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&args)
	if err != nil {
		return err
	}
	return nil
}

func (args *DeleteUserRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(args)
	if err != nil {
		return err
	}
	return nil
}
