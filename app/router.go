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

	// Book part
	bkRepo := repo.NewBookRepo(db)
	bkService := service.NewBookService(bkRepo)
	bkController := controller.NewBookController(bkService)

	// User part
	urRepo := repo.NewUserRepo(db)
	urService := service.NewUserService(urRepo)
	urController := controller.NewUserController(urService)

	// Author part
	authRepo := repo.NewAuthorRepo(db)
	authService := service.NewAuthorService(authRepo)
	authController := controller.NewAuthorController(authService)

	r.Route("/", func(r chi.Router) {
		r.Get("/hello", api.ExampleHamdler)
	})

	// book
	r.Route("/book", func(r chi.Router) {
		r.Post("/create", bkController.CreateBook)
		r.Get("/", bkController.GetallBookDetails)
		r.Get("/{id}", bkController.GetBookById)
		r.Put("/update/{id}", bkController.UpdateBook)
		r.Delete("/remove/{id}", bkController.DeleteBookById)
	})

	// user
	r.Route("/user", func(r chi.Router) {
		r.Post("/create", urController.CreateUser)
		r.Get("/", urController.GetallUserDetails)
		r.Get("/{id}", urController.GetUserById)
		r.Put("/update/{id}", urController.UpdateUser)
		r.Delete("/remove/{id}", urController.DeleteUserById)
	})

	// 	// author
	r.Route("/author", func(r chi.Router) {
		r.Post("/create", authController.CreateAuthor)
		r.Get("/{id}", authController.GetAuthorById)
		r.Get("/", authController.GetallAuthorInfo)
		r.Put("/update/{id}", authController.Updateauthor)
		r.Delete("/remove/{id}", authController.DeleteAuthorById)
	})

	return r
}
