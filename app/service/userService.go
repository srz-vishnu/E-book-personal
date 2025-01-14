package service

import (
	"e-book/app/dto"
	"e-book/app/repo"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
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

	err := args.Parse(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse request body: %w", err)
	}

	err = args.Validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate request body: %w", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	userID, err := s.userRepo.CreateUser(args)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &dto.CreateUSerResponse{
		UserId: userID,
	}, nil
}

func (s *userServiceImpl) UpdateUser(r *http.Request) error {

	args := &dto.UpdateUserRequest{}

	err := args.Parse(r.Body)
	if err != nil {
		return fmt.Errorf("failed to parse the requst: %w", err)
	}

	err = args.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate request body :%w", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	err = s.userRepo.UpdateUser(args, args.UserId)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (s *userServiceImpl) GetUserById(r *http.Request) (*dto.GetUserDetailResponse, error) {
	args := &dto.GetUserDetailRequest{}

	err := args.Parse(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the requst: %w", err)
	}

	err = args.Validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate the requst: %w", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	var userName, userMail string

	userName, userMail, err = s.userRepo.GetOneUser(args.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to get the user details: %w", err)
	}

	return &dto.GetUserDetailResponse{
		UserName: userName,
		Mail:     userMail,
	}, nil
}

func (s *userServiceImpl) GetallUserDetails(r *http.Request) ([]*dto.UserDetails, error) {

	allUserDetails, err := s.userRepo.GetAllUsers()
	if err != nil {
		fmt.Printf(" error while getting all user details %v", err)
		return nil, err
	}

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

	err := args.Parse(r.Body)
	if err != nil {
		return fmt.Errorf("failed to parse the requst: %w", err)
	}

	err = args.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate the requst: %w", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	err = s.userRepo.DeleteUserId(args.UserId)
	if err != nil {
		return fmt.Errorf("failed to delete the user : %w", err)
	}
	log.Info().Msg("succesfully deleted user")

	return nil
}
