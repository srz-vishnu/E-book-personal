package repo

import (
	"errors"
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
	// slice of User structs to hold the results
	var books []Books

	// Fetching the authorname fields from all authors
	result := db.Model(&Books{}).Select("name").Find(&books)

	// Check for any errors during execution
	if result.Error != nil {
		return nil, result.Error
	}

	// Slices to store the retrieved  authornames
	var authorNames []string

	// Iterate over the retrieved authors and extract name one by one
	for _, author := range authors {
		authorNames = append(authorNames, author.Name)
	}

	return authorNames, nil
}
