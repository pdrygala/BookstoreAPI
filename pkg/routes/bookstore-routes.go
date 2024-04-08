package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"pdrygala.com/go-bookstore-api/pkg/controllers"
)

func NewServer() http.Handler {
	r := chi.NewRouter()

	// Defines routes
	r.Route("/book/", func(r chi.Router) {
		r.Post("/", controllers.CreateBook)
		r.Get("/", controllers.GetBook)
		r.Get("/{bookId}", controllers.GetBookById) // This line captures the bookId parameter correctly
		r.Put("/{bookId}", controllers.UpdateBook)
		r.Delete("/{bookId}", controllers.DeleteBook)
	})

	return r
}
