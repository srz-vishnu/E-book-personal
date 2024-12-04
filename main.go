package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	db := ConnectDb()
	defer db.Close()

	// Insert a new user
	userID, err := CreatUser(db, "sherin@gmail.com", "sherin", "password", "text")
	if err != nil {
		log.Fatal("Error creating user:", err)
	}

	fmt.Printf("New book created with ID: %d\n", userID)

	// // Get all users
	usernames, mails, err := GetAllUsers(db)
	if err != nil {
		log.Fatalf("Error retrieving users: %v", err)
	}

	// Print the usernames and mails
	for i := range usernames {
		fmt.Printf("Username: %s, Mail: %s\n", usernames[i], mails[i])
	}

	// Get one User
	userId := int64(3)
	username, mail, err := GetOneUser(db, userId)
	if err != nil {
		log.Fatalf("Error retrieving user: %v", err)
	}

	// Print the result
	fmt.Printf("User found with id: %d\n Username: %s\n Mail: %s\n", userId, username, mail)

	// Deleting an user with id given
	err = DeleteUser(db, 3) // Delete the user with ID 5
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	} else {
		fmt.Println("User with id 3 deleted successfully")
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

func GetOneUser(db *sql.DB, userId int64) (string, string, error) {
	query := `SELECT mail, username FROM users WHERE id = $1`

	// variable to store usernames and mails
	var username, mail string

	// Use QueryRow since you're expecting a single row result
	//err = db.QueryRow(query, userId).Scan(&mail, &username)
	err := db.QueryRow(query, userId).Scan(&mail, &username)

	if err != nil {
		return "", "", err
	}

	return username, mail, nil
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

func DeleteUser(db *sql.DB, userID int64) error {
	// SQL query to delete a user by ID
	query := `UPDATE users SET is_deleted=$1, deleted_at=$2 WHERE id=$3`

	// Execute the query
	result, err := db.Exec(query, true, time.Now().UTC(), userID)
	if err != nil {
		return err
	}

	// Check if any rows affected.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with ID %d", userID)
	}

	return nil
}
