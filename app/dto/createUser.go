package dto

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

// struct to hold the data for creating a user
type CreateUserRequest struct {
	UserId   int64  `json:"userid"`
	Mail     string `json:"mail"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

type CreateUSerResponse struct {
	UserId int64 `json:"userid"`
}

func (args *CreateUserRequest) Parse(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
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
