package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Postgres driver import
)

func main() {
	db := ConnectDb()
	defer db.Close()

	// Insert a new user
	// userID, err := CreatUser(db, "anandhu@gmail.com", "anandhu", "password", "234")
	// if err != nil {
	// 	log.Fatal("Error creating user:", err)
	// }

	// fmt.Printf("New book created with ID: %d\n", userID)

	// Insert a new author
	// authorName, ID, err := CreateAuthor(db, "vishnu")
	// if err != nil {
	// 	log.Fatal("error creating author:", err)
	// }
	// fmt.Printf("new author is created with ID: %d and Name: %s", ID, authorName)

	// Insert a new book
	// bookID, err := CreateBook(db, "noble", "fictional", 3, 7, 1)
	// if err != nil {
	// 	log.Fatal("Error creating book:", err)
	// }

	// fmt.Printf("New book created with ID: %d\n", bookID)

	// Get all users
	usernames, mails, err := GetAllUsers(db)
	if err != nil {
		log.Fatalf("Error retrieving users: %v", err)
	}

	// Print the usernames and mails
	for i := range usernames {
		fmt.Printf("Username: %s, Mail: %s\n", usernames[i], mails[i])
	}

}

func ConnectDb() *sql.DB {
	// PostgreSQL connection string
	connStr := "user=postgres password=1234 dbname=CrudApp sslmode=disable"

	var err error

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v\n", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}

	fmt.Println("Successfully connected to the database!")

	return db
}

type User struct {
	ID       int64  `json:"id"`
	Mail     string `json:"mail"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

func CreatUser(db *sql.DB, mail string, username string, password string, salt string) (int64, error) {
	query := `
	INSERT INTO users (mail, username, password, salt)
	VALUES ($1, $2, $3, $4) RETURNING id`

	var identity int64
	err := db.QueryRow(query, mail, username, password, salt).Scan(&identity)
	if err != nil {
		return 0, err
	}

	return identity, nil
}

func GetAllUsers(db *sql.DB) ([]string, []string, error) {
	query := `SELECT mail, username FROM users`

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	// Slices to store usernames and mails
	var userNames []string
	var mails []string

	// Iterate over rows
	for rows.Next() {
		var mail, username string
		// Scan the values into mail and username variables
		err := rows.Scan(&mail, &username)
		if err != nil {
			return nil, nil, err
		}

		// Append the retrieved values to the slices
		mails = append(mails, mail)
		userNames = append(userNames, username)
	}

	return userNames, mails, nil
}

func CreateAuthor(db *sql.DB, name string) (string, int64, error) {
	query := `INSERT INTO authors (name) VALUES ($1) RETURNING name, id`

	var Name string
	var ID int64
	err := db.QueryRow(query, name).Scan(&Name, &ID)
	if err != nil {
		return "", 0, err
	}
	return Name, ID, nil
}

func CreateBook(db *sql.DB, title, content string, createdBy int64, authorID, status int64) (int64, error) {
	query := `INSERT INTO BOOKS (title, content, created_by, author_id, status) VALUES ($1,$2,$3,$4,$5) RETURNING id`

	var ID int64
	err := db.QueryRow(query, title, content, createdBy, authorID, status).Scan(&ID)
	if err != nil {
		return 0, err
	}
	return ID, nil
}
