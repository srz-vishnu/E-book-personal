package repo

import (
	"e-book/app/dto"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BookRepo interface {
	CreateBook(args *dto.BookInputRequest) (int64, error)
	UpdateBook(args *dto.BookUpdateRequest) error
	GetAllBooks() ([]Book, error)
	GetOneBook(bookId int) (string, error)
	DeleteBookById(args *dto.BookDeleteRequest) error
}

type bookRepoImpl struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) BookRepo {
	return &bookRepoImpl{
		db: db,
	}
}

type Book struct {
	ID        int64      `gorm:"primaryKey;autoIncrement:true"`
	Title     string     `gorm:"type:varchar(255);not null"`
	AuthorID  int64      `gorm:"not null"`
	Content   string     `gorm:"type:text;not null"`
	CreatedAt time.Time  `gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	Status    int        `gorm:"type:int;not null;check:status in (1, 2, 3)"`
	DeletedBy *int       `gorm:"index"`
	DeletedAt *time.Time `gorm:"type:timestamp with time zone"`
	CreatedBy int64      `gorm:"not null"`
	UpdatedBy *int       `gorm:"index"`
}

func (r *bookRepoImpl) CreateBook(args *dto.BookInputRequest) (int64, error) {
	book := Book{
		Title:     args.Title,
		AuthorID:  args.AuthorID,
		Content:   args.Content,
		CreatedBy: args.CreatedBy,
		Status:    1,
		DeletedBy: nil,
		DeletedAt: nil,
		UpdatedBy: nil,
	}

	//GORM's Create method to insert the new book
	if err := r.db.Table("books").Create(book).Error; err != nil {
		return 0, err
	}

	return book.ID, nil
}

func (r *bookRepoImpl) GetOneBook(bookId int) (string, error) {

	// struct to hold the result
	var result struct {
		Status int64  `gorm:"column:status"`
		Title  string `gorm:"column:title"`
	}

	// Use GORM's First method to retrieve a single book with id
	err := r.db.Table("books").Select("title, status").Where("id = ?", bookId).First(&result).Error
	if err != nil {
		// Checking error is "record not found"
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", gorm.ErrRecordNotFound // author not found case
		}
		return "", err // other errors
	}

	if result.Status != 1 {
		return "Book is not active (either deleted or pending).", nil
	}

	return result.Title, nil
}

func (r *bookRepoImpl) GetAllBooks() ([]Book, error) {
	var books []Book

	// Fetch all books
	result := r.db.Find(&books)
	if result.Error != nil {
		fmt.Println("error retriving all book details")
		return nil, result.Error
	}

	return books, nil
}

func (r *bookRepoImpl) DeleteBookById(args *dto.BookDeleteRequest) error {

	var book Book
	err := r.db.First(&book, args.BookId).Error
	if err != nil {
		// Return error if the book does not exist
		return err
	}

	// Checking if the book is already deleted or  not
	if book.Status == 3 {
		return fmt.Errorf("the book with ID %d is already deleted", args.BookId)
	}

	// map to hold the fields to update
	updates := map[string]interface{}{
		"deleted_at": time.Now(),
		"deleted_by": args.UserID,
		"status":     3,
	}

	// soft delete so we can keep data for future use
	err = r.db.Model(&Book{}).Where("id = ?", args.BookId).Updates(updates).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *bookRepoImpl) UpdateBook(args *dto.BookUpdateRequest) error {

	updates := map[string]interface{}{
		"title":      args.Title,
		"content":    args.Content,
		"status":     args.Status,
		"updated_by": args.UserID,
	}

	result := r.db.Table("books").Where("id=?", args.ID).Updates(updates)

	if result.Error != nil {
		return result.Error
	}
	// Check if any rows were updated
	if result.RowsAffected == 0 {
		return fmt.Errorf("no active book found with ID %d to update \n", args.ID)
	}

	fmt.Println("book details updated successfully..")
	return nil
}
