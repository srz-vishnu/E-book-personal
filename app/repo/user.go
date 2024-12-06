package repo

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64      `gorm:"primaryKey"`
	Mail      string     `gorm:"column:mail;unique;not null"`
	Username  string     `gorm:"column:username;unique;not null"`
	Password  string     `gorm:"column:password;not null"`
	Salt      string     `gorm:"column:salt;not null"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time  `gorm:"column:updated_at;autoUpdateTime"`
	IsDeleted bool       `gorm:"column:is_deleted;default:false"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	DeletedBy *int64     `gorm:"column:deleted_by"`
}

func CreateUser(db *gorm.DB, mail, username, password, salt string) (int64, error) {
	user := User{
		Mail:     mail,
		Username: username,
		Password: password,
		Salt:     salt,
	}

	// Use GORM's Create method to insert the new user
	if err := db.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}

func GetOneUser(db *gorm.DB, userId int64) (string, string, error) {

	// struct to hold the result
	var result struct {
		Mail     string
		Username string
	}

	// Use GORM's First method to retrieve a single record
	err := db.Table("users").Select("mail, username").Where("id = ?", userId).First(&result).Error
	if err != nil {
		return "", "", err
	}

	return result.Username, result.Mail, nil
}

func GetAllUsers(db *gorm.DB) ([]string, []string, error) {
	// slice of User structs to hold the results
	var users []User

	// Fetching the mail and username fields from all users
	result := db.Model(&User{}).Select("mail, username").Where("is_deleted = ?", false).Find(&users)

	// Check for any errors during execution
	if result.Error != nil {
		return nil, nil, result.Error
	}

	// Slices to store the retrieved mails and usernames
	var mails []string
	var userNames []string

	// Iterate over the retrieved users and extract mail and username one by one
	for _, user := range users {
		mails = append(mails, user.Mail)
		userNames = append(userNames, user.Username)
	}

	return userNames, mails, nil
}

func DeleteUser(db *gorm.DB, userID int64) error {
	updates := map[string]interface{}{
		"is_deleted": true,             // true means its deleted
		"deleted_at": time.Now().UTC(), // Set the deletion timestamp time of delete
	}

	// Use GORM to update the fields of the selected user by ID
	result := db.Table("users").Where("id = ?", userID).Updates(updates)

	if result.Error != nil {
		log.Fatalf("error on delete: %v", result.Error)
		return result.Error
	}

	// Check if any rows were affected (i.e., if the user exists)
	if result.RowsAffected == 0 {
		return fmt.Errorf("no user found with ID %d", userID)
	}

	return nil
}
