package service

import (
	"e-book/app/dto"
	"e-book/app/repo"
	"e-book/pkg/e"

	//"e-book/pkg/log"
	"errors"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type BookService interface {
	CreateBookService(r *http.Request) (*dto.BookOutputResponse, error)
	UpdateBook(r *http.Request) error
	DeleteBookById(r *http.Request) error
	GetBookById(r *http.Request) (*dto.BookDetailsByIdResponse, error)
	GetallBookDetails(r *http.Request) ([]dto.BookDetails, error)
}

type bookServiceImpl struct {
	bookRepo repo.BookRepo
}

func NewBookService(bookRepo repo.BookRepo) BookService {
	return bookServiceImpl{
		bookRepo: bookRepo,
	}
}

// BookService handles creating a book and returns BookOutput
func (s bookServiceImpl) CreateBookService(r *http.Request) (*dto.BookOutputResponse, error) {
	args := &dto.BookInputRequest{}

	err := args.Parse(r)
	if err != nil {
		return nil, e.NewError(e.ErrDecodeRequestBody, "error while parsing", err)
	}

	err = args.Validate()
	if err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "error while validating", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	bookID, err := s.bookRepo.CreateBook(args)
	if err != nil {
		return nil, e.NewError(e.ErrExecuteSQL, "failed to create book", err)
	}
	log.Info().Msgf("Successfully created book with id %d:", bookID)

	return &dto.BookOutputResponse{
		Id:    bookID,
		Title: args.Title,
	}, nil
}

func (s bookServiceImpl) UpdateBook(r *http.Request) error {
	args := &dto.BookUpdateRequest{}

	err := args.Parse(r)
	if err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "error while parsing", err)
	}

	err = args.Validate()
	if err != nil {
		return e.NewError(e.ErrValidateRequest, "error while validating", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	err = s.bookRepo.UpdateBook(args)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Msgf("book is not found", err)
			return e.NewError(e.ErrResourceNotFound, "book not found in the table", err)
		}
		return e.NewError(e.ErrExecuteSQL, "Failed to update book", err)
	}
	log.Info().Msgf("Successfully updated book with id %d:", args.ID)

	return nil
}

func (s bookServiceImpl) DeleteBookById(r *http.Request) error {
	args := &dto.BookDeleteRequest{}

	err := args.Parse(r)
	if err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "error while parsing", err)
	}

	err = args.Validate()
	if err != nil {
		return e.NewError(e.ErrValidateRequest, "error while validating", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	err = s.bookRepo.DeleteBookById(args)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewError(e.ErrResourceNotFound, "book not found in the record", err)
		}
		return e.NewError(e.ErrDecodeRequestBody, "error while dleeting the book", err)
	}
	log.Info().Msgf("Successfully deleted book with id of %d", args.BookId)

	return nil
}

func (s bookServiceImpl) GetBookById(r *http.Request) (*dto.BookDetailsByIdResponse, error) {
	args := &dto.BookDetailsById{}

	err := args.Parse(r)
	if err != nil {
		return nil, e.NewError(e.ErrDecodeRequestBody, "error while parsing", err)
	}

	err = args.Validate()
	if err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "error while validating", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	bookTitle, err := s.bookRepo.GetOneBook(args.BookID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.NewError(e.ErrResourceNotFound, "book not in the table", err)
		}
		log.Error().Msgf("error while getting book detail %v", err)
		return nil, e.NewError(e.ErrDecodeRequestBody, "error while dleeting the book", err)
	}
	log.Info().Msgf("successfully retrived book with title %s", bookTitle)

	return &dto.BookDetailsByIdResponse{
		Title: bookTitle,
	}, nil

}

func (s bookServiceImpl) GetallBookDetails(r *http.Request) ([]dto.BookDetails, error) {

	allBookDetails, err := s.bookRepo.GetAllBooks()
	if err != nil {
		return nil, e.NewError(e.ErrExecuteSQL, "error while getting all book details", err)
	}
	log.Info().Msg("Successfully retrived all book details")

	var bookDetails []dto.BookDetails

	for _, book := range allBookDetails {
		bookDetail := dto.BookDetails{
			Title:  book.Title,
			Author: book.AuthorID,
		}
		bookDetails = append(bookDetails, bookDetail)
		fmt.Printf("the book details: %v", bookDetails)
	}

	return bookDetails, nil

}
