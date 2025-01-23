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

type AuthorService interface {
	CreateAuthor(r *http.Request) (*dto.CreateAuthorResponse, error)
	GetAuthorById(r *http.Request) (*dto.GetAuthorDetailResponse, error)
	UpdateAuthor(r *http.Request) (err error)
	GetallAuthorDetails(r *http.Request) ([]dto.AuthorDetail, error)
	DeleteAuthorById(r *http.Request) (err error)
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
		return nil, e.NewError(e.ErrDecodeRequestBody, "error while validating", err)
	}

	err = args.Validate()
	if err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "error while validating", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	authorId, err := s.authorRepo.CreateAuthor(&args)
	if err != nil {
		return nil, e.NewError(e.ErrCreateAuthor, "error while creating author", err)
	}
	log.Info().Msgf("Successfully created author with id %d:", authorId)

	return &dto.CreateAuthorResponse{
		AuthorName: args.AuthorName,
		AuthorId:   authorId,
	}, nil
}

func (s *authorServiceImpl) GetAuthorById(r *http.Request) (*dto.GetAuthorDetailResponse, error) {
	args := dto.GetAuthorDetailRequest{}

	err := args.Parse(r)
	if err != nil {
		return nil, e.NewError(e.ErrDecodeRequestBody, "error while validating", err)
	}

	err = args.Validate()
	if err != nil {
		return nil, e.NewError(e.ErrDecodeRequestBody, "error while validating", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	authorName, err := s.authorRepo.GetOneauthor(args.AuthorId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.NewError(e.ErrResourceNotFound, "author not found in the table", err)
		}
		return nil, e.NewError(e.ErrGetAuthorById, "error while getting author by id", err)
	}

	return &dto.GetAuthorDetailResponse{
		AuthorName: authorName,
		AuthorId:   args.AuthorId,
	}, nil
}

func (s *authorServiceImpl) UpdateAuthor(r *http.Request) error {
	args := dto.AuthorupdateRequest{}

	err := args.Parse(r)
	if err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "error while decoding", err)
	}

	err = args.Validate()
	if err != nil {
		return e.NewError(e.ErrValidateRequest, "error while validating", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	err = s.authorRepo.UpdateAuthor(args.UserID, args.AuthorId, args.AuthorName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewError(e.ErrResourceNotFound, "author not found in the table", err)
		}
		return e.NewError(e.ErrUpdateAuthor, "error while author author by id", err)
	}
	log.Info().Msg("Successfully author profile updated")

	return nil

}

func (s *authorServiceImpl) GetallAuthorDetails(r *http.Request) ([]dto.AuthorDetail, error) {

	allAuthorDet, err := s.authorRepo.GetAllAuthor()
	if err != nil {
		return nil, e.NewError(e.ErrGetAllAuthorDetails, "error while getting all author details", err)
	}
	log.Info().Msg("Successfully author profile updated")

	var authorDetails []dto.AuthorDetail

	for _, author := range allAuthorDet {
		authorDetail := dto.AuthorDetail{
			AuthorName: author.Name,
			AuthorId:   author.ID,
		}
		authorDetails = append(authorDetails, authorDetail)
		fmt.Printf("the book details: %v", authorDetails)
	}

	return authorDetails, nil

}

func (s *authorServiceImpl) DeleteAuthorById(r *http.Request) error {
	args := &dto.AuthorDeleteRequest{}

	err := args.Parse(r)
	if err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "error while decoding", err)
	}

	err = args.Validate()
	if err != nil {
		return e.NewError(e.ErrValidateRequest, "error while validating", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	err = s.authorRepo.DeleteAuthor(args)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewError(e.ErrResourceNotFound, "author not found in the table", err)
		}
		return e.NewError(e.ErrDeleteAuthor, "error while getting all author details", err)
	}
	log.Info().Msg("Successfully removed the author")

	return nil
}
