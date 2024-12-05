package main

import (
	"e-book/app/repo"
	gormdb "e-book/pkg/gorm_db"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := gormdb.ConnectDb()

	repoConn := repo.User{
		Mail:     "dj@gmail.com",
		Username: "Dj",
		Password: "password1",
		Salt:     "random",
	}

	// Create a new user
	userIdentity, err := repo.CreateUser(db, repoConn.Mail, repoConn.Username, repoConn.Password, repoConn.Salt)
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	// Print the created user's ID
	fmt.Printf("User created with ID: %d\n", userIdentity)

	// Get one User
	idOfUser := int64(3)
	usernamee, mail, err := repo.GetOneUser(db, idOfUser)
	if err != nil {
		log.Fatalf("Error retrieving user: %v", err)
	}

	// Print the result
	fmt.Printf("User found with id: %d\n Username: %s\n Mail: %s\n", idOfUser, usernamee, mail)

	// // Get all users
	usernames, mails, err := repo.GetAllUsers(db)
	if err != nil {
		log.Fatalf("Error retrieving users: %v", err)
	}

	// Print the usernames and mails
	for i := range usernames {
		fmt.Printf("Username: %s, Mail: %s\n", usernames[i], mails[i])
	}

	// Deleting an user with id given
	err = repo.DeleteUser(db, 10) // Delete the user with ID
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	} else {
		fmt.Println("User deleted successfully")
	}

}
