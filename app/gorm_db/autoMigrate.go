package gormdb

import (
	"e-book/app/repo"
	"log"

	"gorm.io/gorm"
)

func Automigration(db *gorm.DB) error {
	if err := db.AutoMigrate(&repo.User{}); err != nil {
		log.Fatal("Migration error for user:", err)
	}

	if err := db.AutoMigrate(&repo.Author{}); err != nil {
		log.Fatal("migration failed for author", err)
	}

	if err := db.AutoMigrate(&repo.Book{}); err != nil {
		log.Fatal("migration failed for book", err)
	}

	log.Println("Migration success")
	return nil
}
