package repo

import (
	"e-book/app/dto"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(args *dto.CreateUserRequest) (int64, error)
	GetOneUser(userId int64) (string, string, error)
	GetAllUsers() ([]User, error)
	DeleteUserId(userId int64) error
	UpdateUser(args *dto.UpdateUserRequest, userID int64) error
}

type userRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepoImpl{
		db: db,
	}
}

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

func (r *userRepoImpl) CreateUser(args *dto.CreateUserRequest) (int64, error) {
	user := User{
		Mail:     args.Mail,
		Username: args.UserName,
		Password: args.Password,
		Salt:     args.Salt,
	}

	// Use GORM's Create method to insert the new user
	if err := r.db.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (r *userRepoImpl) GetOneUser(userId int64) (string, string, error) {

	// struct to hold the result
	var result struct {
		Mail     string
		Username string
	}

	// Use GORM's First method to retrieve a single record
	err := r.db.Table("users").Select("mail, username").Where("id = ?", userId).First(&result).Error
	if err != nil {
		return "", "", err
	}

	return result.Username, result.Mail, nil
}

func (r *userRepoImpl) GetAllUsers() ([]User, error) {
	// slice of User structs to hold the results
	var users []User

	// Fetch all users
	result := r.db.Find(&users)
	if result.Error != nil {
		fmt.Println("error retriving all user details")
		return nil, result.Error
	}

	return users, nil
}

func (r *userRepoImpl) UpdateUser(args *dto.UpdateUserRequest, userID int64) error {

	updates := map[string]interface{}{
		"mail":       args.Mail,
		"username":   args.UserName,
		"password":   args.Password,
		"updated_at": time.Now(),
	}

	result := r.db.Table("users").Where("id=?", userID).Updates(updates)

	if result.Error != nil {
		return result.Error
	}
	// Check if any rows were updated
	if result.RowsAffected == 0 {
		return fmt.Errorf("no active user found with ID %d to update", userID)
	}

	log.SetFlags(0)
	log.Println("user password updated successfully..")
	return nil
}

func (r *userRepoImpl) DeleteUserId(userID int64) error {
	updates := map[string]interface{}{
		"is_deleted": true,       // true means its deleted
		"deleted_at": time.Now(), // Set the deletion timestamp time of delete
	}

	// Use GORM to update the fields of the selected user by ID
	result := r.db.Table("users").Where("id = ?", userID).Updates(updates)

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
