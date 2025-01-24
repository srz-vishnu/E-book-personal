package service

import (
	"e-book/app/dto"
	"e-book/app/repo"
	"e-book/pkg/e"
	"errors"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(r *http.Request) (*dto.CreateUSerResponse, error)
	UpdateUser(r *http.Request) error
	GetUserById(r *http.Request) (*dto.GetUserDetailResponse, error)
	GetallUserDetails(r *http.Request) ([]*dto.UserDetails, error)
	DeleteUserById(r *http.Request) error
}

type userServiceImpl struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

func (s *userServiceImpl) CreateUser(r *http.Request) (*dto.CreateUSerResponse, error) {
	args := &dto.CreateUserRequest{}

	err := args.Parse(r)
	if err != nil {
		return nil, e.NewError(e.ErrDecodeRequestBody, "error while parsing", err)
	}

	err = args.Validate()
	if err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "error while validating", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	userID, err := s.userRepo.CreateUser(args)
	if err != nil {
		return nil, e.NewError(e.ErrExecuteSQL, "error while creating user", err)
	}
	log.Info().Msgf("Successfully created user with id %d", userID)

	return &dto.CreateUSerResponse{
		UserId: userID,
	}, nil
}

func (s *userServiceImpl) UpdateUser(r *http.Request) error {

	args := &dto.UpdateUserRequest{}

	err := args.Parse(r)
	if err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "error while parsing", err)
	}

	err = args.Validate()
	if err != nil {
		return e.NewError(e.ErrValidateRequest, "error validating the req body", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	err = s.userRepo.UpdateUser(args, args.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewError(e.ErrResourceNotFound, "user not found in the table", err)
		}
		return e.NewError(e.ErrExecuteSQL, "error while updating the user details", err)
	}

	log.Info().Msg("Successfully completed parsing and validation of request body")

	return nil
}

func (s *userServiceImpl) GetUserById(r *http.Request) (*dto.GetUserDetailResponse, error) {
	args := &dto.GetUserDetailRequest{}

	err := args.Parse(r)
	if err != nil {
		return nil, e.NewError(e.ErrDecodeRequestBody, "error while parsing the req body", err)
	}

	err = args.Validate()
	if err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "error while validating the req body", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	var userName, userMail string

	userName, userMail, err = s.userRepo.GetOneUser(args.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.NewError(e.ErrResourceNotFound, "user not found in the table", err)
		}
		return nil, e.NewError(e.ErrExecuteSQL, "error while retryving given user by id", err)
	}
	log.Info().Msgf("Successfully retrived given user details")

	return &dto.GetUserDetailResponse{
		UserName: userName,
		Mail:     userMail,
	}, nil
}

func (s *userServiceImpl) GetallUserDetails(r *http.Request) ([]*dto.UserDetails, error) {

	allUserDetails, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, e.NewError(e.ErrDecodeRequestBody, "error while parsing the req.body", err)
	}
	log.Info().Msgf("Successfully retrived all user details", allUserDetails)

	var userDetails []*dto.UserDetails

	for _, user := range allUserDetails {
		userDetail := &dto.UserDetails{
			UserName: user.Username,
			Mail:     user.Mail,
		}
		userDetails = append(userDetails, userDetail)
		fmt.Printf("all user details: %v", userDetails)
	}

	return userDetails, nil
}

func (s *userServiceImpl) DeleteUserById(r *http.Request) error {
	args := dto.DeleteUserRequest{}

	err := args.Parse(r)
	if err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "error while parsing the request body", err)
	}

	err = args.Validate()
	if err != nil {
		return e.NewError(e.ErrValidateRequest, "error while validating req body", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	err = s.userRepo.DeleteUserId(args.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewError(e.ErrResourceNotFound, "user not found in the table", err)
		}
		return e.NewError(e.ErrExecuteSQL, "error while deleting the user", err)
	}
	log.Info().Msgf("succesfully deleted user with id %d", args.UserId)

	return nil
}
