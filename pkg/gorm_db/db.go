package gormdb

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	user     = "postgres"
	password = "1234"
	host     = "localhost"
	port     = 5432
	dbname   = "CrudApp"
)

func ConnectDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", user, password, host, port, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connection error", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal("connection error", err)
	}

	// Test the connection
	err = sqlDb.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}

	fmt.Println("Successfully connected to the database!!!")

	return db, nil
}
