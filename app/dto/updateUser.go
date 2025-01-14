package dto

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type UpdateUserRequest struct {
	UserId   int64  `json:"userid"`
	Mail     string `json:"mail" validate:"required"`
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserResponse struct {
	Mail     string `json:"mail"`
	UserName string `json:"username"`
}

func (args *UpdateUserRequest) Parse(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&args)
	if err != nil {
		return err
	}
	return nil
}

func (args *UpdateUserRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(args)
	if err != nil {
		return err
	}
	return nil
}
