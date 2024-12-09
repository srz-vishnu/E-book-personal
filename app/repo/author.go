package repo

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        int64     `gorm:"primaryKey;autoIncrement:true"`
	Name      string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone;default:current_timestamp"`
	CreatedBy int64     `gorm:"type:int"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone;default:current_timestamp"`
	UpdatedBy int64     `gorm:"type:int"`
	DeletedAt time.Time `gorm:"type:timestamp with time zone;default:null"`
	DeletedBy *int64    `gorm:"type:int"`
	Status    bool      `gorm:"type:boolean;default:true"`
}

func CreateAuthor(db *gorm.DB, name string, createdBy, updatedBy int64) (int64, error) {
	author := Author{
		Name:      name,
		CreatedBy: createdBy,
		Status:    true, //default status true
		UpdatedBy: updatedBy,
		DeletedBy: nil,
	}

	//GORM's Create method to insert the new author
	if err := db.Create(&author).Error; err != nil {
		return 0, err
	}

	return author.ID, nil
}

func GetOneauthor(db *gorm.DB, userId int64) (string, error) {

	// struct to hold the result
	var result struct {
		Name string `gorm:"column:name"`
	}

	// Use GORM's First method to retrieve a single record
	//err := //db.Table("authors").Select("name").Where("id = ?", userId).First(&result).Error
	err := db.Table("authors").Select("name").Where("id = ? AND status = ?", userId, true).First(&result).Error
	if err != nil {
		// Checking error is "record not found"
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil // author not found case
		}
		return "", err // other errors
	}

	return result.Name, nil
}

func GetAllAuthor(db *gorm.DB) ([]string, error) {
	// slice of User structs to hold the results
	var authors []Author

	// Fetching the authorname fields from all authors
	result := db.Model(&Author{}).Select("name").Find(&authors)

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

func DeleteAuthor(db *gorm.DB, authorID int64, deletedBy int64) error {
	// map to hold the fields to update
	updates := map[string]interface{}{
		"deleted_at": time.Now(),
		"deleted_by": deletedBy,
		"status":     false,
	}

	// soft delete so we can keep data for future use
	err := db.Model(&Author{}).Where("id = ? AND status = ?", authorID, true).Updates(updates).Error
	if err != nil {
		return err
	}

	return nil
}
