package controller

import (
	"e-book/app/service"
	"e-book/pkg/api"
	"e-book/pkg/e"
	"net/http"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	GetallUserDetails(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUserById(w http.ResponseWriter, r *http.Request)
}

type userControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userControllerImpl{
		userService: userService,
	}
}

func (c *userControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	resp, err := c.userService.CreateUser(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "failed to create user")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, resp)
}

func (c *userControllerImpl) GetUserById(w http.ResponseWriter, r *http.Request) {
	resp, err := c.userService.GetUserById(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "failed to get user detail")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, resp)
}

func (c *userControllerImpl) GetallUserDetails(w http.ResponseWriter, r *http.Request) {
	resp, err := c.userService.GetallUserDetails(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "failed to get all user details")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, resp)
}

func (c *userControllerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	err := c.userService.UpdateUser(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "user updation failed")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "success")
}

func (c *userControllerImpl) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	err := c.userService.DeleteUserById(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "failed to delete user")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "success")
}
