package cmd

import (
	"log"

	"e-book/app"
	"e-book/pkg/api"

	gormdb "e-book/app/gorm_db"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Root short description",
	Long:  "Root long description",
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Root short description",
	Long:  "Root long description",
	Run:   startAPI,
}

func startAPI(*cobra.Command, []string) {
	db, err := gormdb.ConnectDb() // Establish DB connection
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}

	r := app.APIRouter(db)
	api.Start(r)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
