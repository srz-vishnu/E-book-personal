package dto

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

type GetUserDetailRequest struct {
	UserId int64 `json:"userid" validate:"required"`
}

type GetUserDetailResponse struct {
	UserName string `json:"username "`
	Mail     string `json:"mail"`
}

func (args *GetUserDetailRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	log.Printf("Extracted ID from URL issssssssssssssssss: '%s'\n", strID)
	if strID == "" {
		return fmt.Errorf("id parameter is missing or empty")
	}
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	args.UserId = int64(intID)
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
