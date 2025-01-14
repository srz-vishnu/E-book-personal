package dto

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type GetUserDetailRequest struct {
	UserId int64 `json:"userid" validate:"required"`
}

type GetUserDetailResponse struct {
	UserName string `json:"username "`
	Mail     string `json:"mail"`
}

func (args *GetUserDetailRequest) Parse(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&args)
	if err != nil {
		return err
	}
	return nil
}

func (args *GetUserDetailRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(args)
	if err != nil {
		return err
	}
	return nil
}
