package dto

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type GetAuthorDetailRequest struct {
	AuthorId int64 `json:"authorid"`
}

type GetAuthorDetailResponse struct {
	AuthorName string `json:"authorname"`
	AuthorId   int64  `json:"authorid"`
}

func (args *GetAuthorDetailRequest) Parse(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&args)
	if err != nil {
		return err
	}
	return nil
}

func (args *GetAuthorDetailRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(args)
	if err != nil {
		return err
	}
	return nil
}
