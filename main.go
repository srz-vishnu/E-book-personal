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

	userConn := repo.User{
		Mail:     "test1@gmail.com",
		Username: "testUser1",
		Password: "password1",
		Salt:     "random",
	}

	// Create a new user
	userIdentity, err := repo.CreateUser(db, userConn.Mail, userConn.Username, userConn.Password, userConn.Salt)
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

	// Get all users
	usernames, mails, err := repo.GetAllUsers(db)
	if err != nil {
		log.Fatalf("Error retrieving users: %v", err)
	}

	//Print the usernames and mails
	for i := range usernames {
		fmt.Printf("Username: %s, Mail: %s\n", usernames[i], mails[i])
	}

	// Deleting an user with id given
	err = repo.DeleteUser(db, 12) // Delete the user with ID
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	} else {
		fmt.Println("User deleted successfully")
	}

	// Updating user password

	userID := int64(12)
	newPassword := "test@34"
	err = repo.UpdateUser(db, userID, newPassword)
	if err != nil {
		log.Fatalf("err in updating user %v", err)
	}

	fmt.Printf("User with id %d successfully updated password", userID)

	// Author part

	authorConn := repo.Author{
		Name:      "abhi",
		CreatedBy: 2, //userId
		UpdatedBy: 2,
	}

	// Create a new author
	authorIdentity, err := repo.CreateAuthor(db, authorConn.Name, authorConn.CreatedBy, authorConn.UpdatedBy)
	if err != nil {
		log.Fatalf("Error creating author: %v", err)
	}

	//Print the created author's ID
	fmt.Printf("Author created with ID: %d\n", authorIdentity)

	// Get one Author
	idOfAuthor := int64(18)
	authorName, err := repo.GetOneauthor(db, idOfAuthor)
	if err != nil {
		log.Fatalf("Error retrieving author: %v", err)
	}

	// Print the result
	if authorName == "" {
		fmt.Printf("No author found with id: %d, its already deleted \n", idOfAuthor)
	} else {
		fmt.Printf("Author found with id: %d\n Authorname: %s\n", idOfAuthor, authorName)
	}

	// Get all authors
	authorNames, err := repo.GetAllAuthor(db)
	if err != nil {
		log.Fatalf("Error retrieving all authors: %v", err)
	}

	// Print the author names
	for i := range authorNames {
		fmt.Printf("Author name: %s \n", authorNames[i])
	}

	authorIdToDelete := int64(17)
	deletedBy := int64(3)
	err = repo.DeleteAuthor(db, authorIdToDelete, deletedBy)
	if err != nil {
		log.Fatalf("Error deleting author: %v", err)
	} else {
		fmt.Printf("Author with id %d successfully deleted\n", authorIdToDelete)
	}

	// Updating author details
	userId := int64(12)
	authorId := int64(23)
	newName := "raina"
	err = repo.UpdateAuthor(db, userId, authorId, newName)
	if err != nil {
		fmt.Printf("error while updating author details %v", err)
	}

	fmt.Println("author details updated for author with id:", authorId)

	// book part

	// Creating a book
	bookId, err := repo.CreateBook(db, "physcology of book", 17, 2, 1, "motivational")
	if err != nil {
		log.Fatalf("Error Creating a book: %v", err)
	}

	// Print the created Book's ID
	fmt.Printf("Book created with ID: %d\n", bookId)

	//Get one book
	bookTitle, err := repo.GetOneBook(db, 12, 1)
	if err != nil {
		log.Fatalf("error retriving a single book", err)
	}

	// Print the created book's title
	if bookTitle == "" {
		fmt.Println("No book found with the given ID and status.")
	} else {
		// Print the book's title
		fmt.Printf("Book found with title: %s\n", bookTitle)
	}

	// Get all Books
	bookTitles, err := repo.GetAllBooks(db)
	if err != nil {
		log.Fatalf("Error retrieving all books: %v", err)
	}

	// Print the book title
	for i := range bookTitle {
		fmt.Printf("book title: %s \n", bookTitles[i])
	}

	// Delete book
	deleteBookByID := int64(12)
	UserId := int64(3)

	err = repo.DeleteBookById(db, int64(deleteBookByID), userID)
	if err != nil {
		log.Fatalf("error while deleting book: %v", err)
	}

	fmt.Printf("book deleted succesfully by user %d\n", UserId)

	// book details updation
	userIdentity = int64(3)
	bookId = int64(12)
	title := "goat life"
	content := "biopic, adventures"
	status := int64(1)
	err = repo.UpdateBook(db, userIdentity, bookId, title, content, status)
	if err != nil {
		fmt.Printf("error while updating author details %v \n", err)
	}

	fmt.Println("book details updated for book with id:", bookId)

}
