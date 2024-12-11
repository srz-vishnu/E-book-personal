package repo

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

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

func CreateBook(db *gorm.DB, title string, authorID, createdBy, status int64, content string) (int64, error) {
	book := Book{
		Title:     title,
		AuthorID:  authorID,
		Content:   content,
		CreatedBy: createdBy,
		Status:    1,
		DeletedBy: nil,
		DeletedAt: nil,
		UpdatedBy: nil,
	}

	// Create method to insert the new book
	if err := db.Create(&book).Error; err != nil {
		return 0, err
	}

	return book.ID, nil
}

func GetOneBook(db *gorm.DB, bookId, status int64) (string, error) {

	// struct to hold the result
	var result struct {
		Status int64  `gorm:"column:status"`
		Title  string `gorm:"column:title"`
	}

	// Use GORM's First method to retrieve a single book with id
	err := db.Table("books").Select("title").Where("id = ? AND status = ?", bookId, status).First(&result).Error
	if err != nil {
		// Checking error is "record not found"
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil // author not found case
		}
		return "", err // other errors
	}

	if result.Status != 1 {
		return "Book is not active (either deleted or pending).", nil
	}

	return result.Title, nil
}

func GetAllBooks(db *gorm.DB) ([]string, error) {
	// slice of book structs to hold the results
	var books []Book

	// Fetching the book title fields from all books
	result := db.Model(&Book{}).Select("title").Find(&books)

	// Check for any errors during execution
	if result.Error != nil {
		return nil, result.Error
	}

	// Slices to store the retrieved  book titles
	var bookTitles []string

	// Iterate over the retrieved books and extract title one by one
	for _, book := range books {
		bookTitles = append(bookTitles, book.Title)
	}

	return bookTitles, nil
}

func DeleteBookById(db *gorm.DB, ID int64, deletedBy int64) error {

	var book Book
	err := db.First(&book, ID).Error
	if err != nil {
		// Return error if the book does not exist
		return err
	}

	// Checking if the book is already deleted or  not
	if book.Status == 3 {
		return fmt.Errorf("the book with ID %d is already deleted", ID)
	}

	// map to hold the fields to update
	updates := map[string]interface{}{
		"deleted_at": time.Now(),
		"deleted_by": deletedBy,
		"status":     3,
	}

	// soft delete so we can keep data for future use
	err = db.Model(&Book{}).Where("id = ?", ID).Updates(updates).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateBook(db *gorm.DB, userId, bookId int64, title, content string, status int64) error {

	updates := map[string]interface{}{
		"title":      title,
		"content":    content,
		"status":     status,
		"updated_by": userId,
	}

	result := db.Table("books").Where("id=?", bookId).Updates(updates)

	if result.Error != nil {
		return result.Error
	}
	// Check if any rows were updated
	if result.RowsAffected == 0 {
		return fmt.Errorf("no active book found with ID %d to update \n", bookId)
	}

	fmt.Println("book details updated successfully..")
	return nil
}
