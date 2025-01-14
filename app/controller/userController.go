package controller

import (
	"e-book/app/service"
	"fmt"
	"net/http"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
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
	res, err := c.userService.CreateUser(r)

	if err != nil {
		fmt.Println("not sucess", err)
		return
	}
	fmt.Println("response from controller", res)
}
