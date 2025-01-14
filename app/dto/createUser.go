package dto

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

// struct to hold the data for creating a user
type CreateUserRequest struct {
	Mail     string `json:"mail"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

type CreateUSerResponse struct {
	UserId int64 `json:"userid"`
}

func (args *CreateUserRequest) Parse(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&args)
	if err != nil {
		return err
	}
	return nil
}

func (args *CreateUserRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(args)
	if err != nil {
		return err
	}
	return nil
}
