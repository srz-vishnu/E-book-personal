package dto

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

type DeleteUserRequest struct {
	UserId int64 `json:"userid" validate:"required"`
}

func (args *DeleteUserRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	args.UserId = int64(intID)

	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
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
