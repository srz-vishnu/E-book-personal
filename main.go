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

	// userConn := repo.User{
	// 	Mail:     "test1@gmail.com",
	// 	Username: "testUser1",
	// 	Password: "password1",
	// 	Salt:     "random",
	// }

	// // Create a new user
	// userIdentity, err := repo.CreateUser(db, userConn.Mail, userConn.Username, userConn.Password, userConn.Salt)
	// if err != nil {
	// 	log.Fatalf("Error creating user: %v", err)
	// }

	// // Print the created user's ID
	// fmt.Printf("User created with ID: %d\n", userIdentity)

	// Get one User
	// idOfUser := int64(3)
	// usernamee, mail, err := repo.GetOneUser(db, idOfUser)
	// if err != nil {
	// 	log.Fatalf("Error retrieving user: %v", err)
	// }

	// // Print the result
	// fmt.Printf("User found with id: %d\n Username: %s\n Mail: %s\n", idOfUser, usernamee, mail)

	// // Get all users
	// usernames, mails, err := repo.GetAllUsers(db)
	// if err != nil {
	// 	log.Fatalf("Error retrieving users: %v", err)
	// }

	// Print the usernames and mails
	// for i := range usernames {
	// 	fmt.Printf("Username: %s, Mail: %s\n", usernames[i], mails[i])
	// }

	// Deleting an user with id given
	// err = repo.DeleteUser(db, 10) // Delete the user with ID
	// if err != nil {
	// 	log.Fatalf("Error deleting user: %v", err)
	// } else {
	// 	fmt.Println("User deleted successfully")
	// }

	// Author part

	// authorConn := repo.Author{
	// 	Name:      "ajay",
	// 	CreatedBy: 2,
	// 	UpdatedBy: 2,
	// }

	// // Create a new author
	// authorIdentity, err := repo.CreateAuthor(db, authorConn.Name, authorConn.CreatedBy, authorConn.UpdatedBy)
	// if err != nil {
	// 	log.Fatalf("Error creating author: %v", err)
	// }

	// // Print the created author's ID
	// fmt.Printf("Author created with ID: %d\n", authorIdentity)

	// Get one Author
	// idOfAuthor := int64(15)
	// authorName, err := repo.GetOneauthor(db, idOfAuthor)
	// if err != nil {
	// 	log.Fatalf("Error retrieving author: %v", err)
	// }

	// // Print the result
	// fmt.Printf("Author found with id: %d\n Authorname: %s\n", idOfAuthor, authorName)

	// // Get all authors
	authorNames, err := repo.GetAllAuthor(db)
	if err != nil {
		log.Fatalf("Error retrieving all authors: %v", err)
	}

	// Print the author names
	for i := range authorNames {
		fmt.Printf("Author name: %s \n", authorNames[i])
	}

	authorIdToDelete := int64(15)
	deletedBy := int64(3)
	err = repo.DeleteAuthor(db, authorIdToDelete, deletedBy)
	if err != nil {
		log.Fatalf("Error deleting author: %v", err)
	} else {
		fmt.Printf("Author with id %d successfully deleted\n", authorIdToDelete)
	}

}
