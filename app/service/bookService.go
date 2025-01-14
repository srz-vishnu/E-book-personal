package service

import (
	"e-book/app/dto"
	"e-book/app/repo"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type BookService interface {
	CreateBookService(r *http.Request) (*dto.BookOutputResponse, error)
	UpdateBook(r *http.Request) error
	DeleteBookById(r *http.Request) error
	GetBookById(r *http.Request, book dto.BookDetailsById) (*dto.BookDetailsByIdResponse, error)
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

	err := args.Parse(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse request body: %w", err)
	}

	err = args.Validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate request body: %w", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	bookID, err := s.bookRepo.CreateBook(args)
	if err != nil {
		return nil, fmt.Errorf("failed to create book: %w", err)
	}

	return &dto.BookOutputResponse{
		Id:    bookID,
		Title: args.Title,
	}, nil
}

func (s bookServiceImpl) UpdateBook(r *http.Request) error {
	args := &dto.BookUpdateRequest{}

	err := args.Parse(r.Body)
	if err != nil {
		return fmt.Errorf("failed to parse request body: %w", err)
	}

	err = args.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate request body: %w", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	err = s.bookRepo.UpdateBook(args, args.BookID)
	if err != nil {
		fmt.Printf("error while creating an book %v", err)
	}

	return nil
}

func (s bookServiceImpl) DeleteBookById(r *http.Request) error {
	args := &dto.BookDeleteRequest{}

	err := args.Parse(r.Body)
	if err != nil {
		return fmt.Errorf("failed to parse request body: %w", err)
	}

	err = args.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate request body: %w", err)
	}

	err = s.bookRepo.DeleteBookById(args.BookId, args.UserID)
	if err != nil {
		fmt.Printf("error while deleting book %v", err)
	}
	return nil
}

func (s bookServiceImpl) GetBookById(r *http.Request, book dto.BookDetailsById) (*dto.BookDetailsByIdResponse, error) {
	args := &dto.BookDetailsById{}

	err := args.Parse(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse request body: %w", err)
	}

	err = args.Validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate request body: %w", err)
	}

	bookTitle, err := s.bookRepo.GetOneBook(args.BookID)
	if err != nil {
		fmt.Printf("error getting book by given id %v", err)
		return nil, err
	}

	return &dto.BookDetailsByIdResponse{
		Title: bookTitle,
	}, nil

}

func (s bookServiceImpl) GetallBookDetails(r *http.Request) ([]dto.BookDetails, error) {

	allBookDetails, err := s.bookRepo.GetAllBooks()
	if err != nil {
		fmt.Printf(" error while getting all book details %v", err)
		return nil, err
	}

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
