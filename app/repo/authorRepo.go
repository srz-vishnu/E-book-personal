package repo

import (
	"e-book/app/dto"
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type AuthorRepo interface {
	CreateAuthor(args *dto.CreateAuthorRequest) (int64, error)
	GetOneauthor(userId int64) (string, error)
	GetAllAuthor() ([]Author, error)
	DeleteAuthor(args *dto.AuthorDeleteRequest) error
	UpdateAuthor(userId, authorId int64, name string) error
}

type authorRepoImpl struct {
	db *gorm.DB
}

func NewAuthorRepo(db *gorm.DB) AuthorRepo {
	return &authorRepoImpl{
		db: db,
	}
}

type Author struct {
	ID        int64          `gorm:"primaryKey;autoIncrement:true"`
	Name      string         `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `gorm:"type:timestamp with time zone;default:current_timestamp"`
	CreatedBy int64          `gorm:"type:int"`
	UpdatedAt time.Time      `gorm:"type:timestamp with time zone;default:current_timestamp"`
	UpdatedBy *int64         `gorm:"type:int"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	DeletedBy *int64         `gorm:"type:int"`
	Status    bool           `gorm:"type:boolean;default:true"`
}

func (r *authorRepoImpl) CreateAuthor(args *dto.CreateAuthorRequest) (int64, error) {
	author := &Author{
		Name:      args.AuthorName,
		CreatedBy: args.UserId,
		Status:    true, //default status true
	}

	//GORM's Create method to insert the new author
	if err := r.db.Table("authors").Create(author).Error; err != nil {
		return 0, err
	}

	return author.ID, nil
}

func (r *authorRepoImpl) GetOneauthor(userId int64) (string, error) {

	// struct to hold the result
	var result struct {
		Name string `gorm:"column:name"`
	}

	// Use GORM's First method to retrieve a single record of author
	err := r.db.Table("authors").Select("name").Where("id = ? AND status = ?", userId, true).First(&result).Error
	if err != nil {
		// If the record is not found, return an appropriate error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", gorm.ErrRecordNotFound
		}
		return "", err // other errors
	}

	return result.Name, nil
}

func (r *authorRepoImpl) GetAllAuthor() ([]Author, error) {
	// slice of Author structs to hold the results
	var authors []Author

	// Fetching the id and name fields from all authors
	result := r.db.Model(&Author{}).Select("id, name").Find(&authors)

	// Check for any errors during execution
	if result.Error != nil {
		return nil, result.Error
	}

	return authors, nil
}

func (r *authorRepoImpl) DeleteAuthor(args *dto.AuthorDeleteRequest) error {
	// map to hold the fields to update
	updates := map[string]interface{}{
		"deleted_at": time.Now(),
		"deleted_by": args.UserID,
		"status":     false,
	}

	// soft delete so we can keep data for future use
	err := r.db.Table("authors").Where("id = ? AND status = ?", args.AuthorId, true).Updates(updates).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *authorRepoImpl) UpdateAuthor(userId, authorId int64, name string) error {

	updates := map[string]interface{}{
		"name":       name,
		"updated_at": time.Now(),
		"updated_by": userId,
	}

	result := r.db.Table("authors").Where("id=?", authorId).Updates(updates)

	if result.Error != nil {
		return result.Error
	}
	// Check if any rows were updated
	if result.RowsAffected == 0 {
		return fmt.Errorf("no active author found with ID %d to update ", authorId)
	}

	log.SetFlags(0)
	log.Println("author details updated successfully..")
	return nil
}
