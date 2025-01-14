package service

import (
	"e-book/app/dto"
	"e-book/app/repo"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type AuthorService interface {
	CreateAuthor(r *http.Request) (*dto.CreateAuthorResponse, error)
	GetAuthorById(r *http.Request) (*dto.GetAuthorDetailResponse, error)
	UpdateAuthor(r *http.Request) error
	// GetallAuthorDetails(r *http.Request) ([]*dto.UserDetails, error)
	// DeleteAuthorById(r *http.Request) error
}

type authorServiceImpl struct {
	authorRepo repo.AuthorRepo
}

func NewAuthorService(authorRepo repo.AuthorRepo) AuthorService {
	return &authorServiceImpl{
		authorRepo: authorRepo,
	}
}

func (s *authorServiceImpl) CreateAuthor(r *http.Request) (*dto.CreateAuthorResponse, error) {
	args := dto.CreateAuthorRequest{}

	err := args.Parse(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while parsing :%w", err)
	}

	err = args.Validate()
	if err != nil {
		return nil, fmt.Errorf("error while validating :%w", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	authorId, err := s.authorRepo.CreateAuthor(&args)
	if err != nil {
		return nil, fmt.Errorf("error while creating author :%d", err)
	}

	return &dto.CreateAuthorResponse{
		AuthorName: args.AuthorName,
		AuthorId:   authorId,
	}, nil
}

func (s *authorServiceImpl) GetAuthorById(r *http.Request) (*dto.GetAuthorDetailResponse, error) {
	args := dto.GetAuthorDetailRequest{}

	err := args.Parse(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while parsing :%w", err)
	}

	err = args.Validate()
	if err != nil {
		return nil, fmt.Errorf("error while validating :%w", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	authorName, err := s.authorRepo.GetOneauthor(args.AuthorId)
	if err != nil {
		return nil, fmt.Errorf("error while creating author by given userId :%d", err)
	}

	return &dto.GetAuthorDetailResponse{
		AuthorName: authorName,
		AuthorId:   args.AuthorId,
	}, nil
}

func (s *authorServiceImpl) UpdateAuthor(r *http.Request) error {
	args := dto.AuthorupdateRequest{}

	err := args.Parse(r.Body)
	if err != nil {
		return fmt.Errorf("error while parsing :%w", err)
	}

	err = args.Validate()
	if err != nil {
		return fmt.Errorf("error while validating :%w", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	err = s.authorRepo.UpdateAuthor(args.UserID, args.AuthorId, args.AuthorName)
	if err != nil {
		return fmt.Errorf("error while updating author :%d", err)
	}
	log.Info().Msg("Successfully updated")

	return nil

}
