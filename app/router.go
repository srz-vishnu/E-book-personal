package app

import (
	"e-book/app/controller"
	"e-book/app/repo"
	"e-book/app/service"
	api "e-book/pkg/api"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func APIRouter(db *gorm.DB) chi.Router {
	r := chi.NewRouter()

	bkRepo := repo.NewBookRepo(db)
	bkService := service.NewBookService(bkRepo)
	bkController := controller.NewBookController(bkService)

	r.Route("/", func(r chi.Router) {
		r.Get("/hello", api.ExampleHamdler)

		r.Post("/create/book", bkController.CreateBook)

		//r.Get("/create/Book",)
	})

	return r
}
