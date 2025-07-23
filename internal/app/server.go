package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/SamJohn04/notes-backend/internal/config"
	"github.com/SamJohn04/notes-backend/internal/handler"
	"github.com/SamJohn04/notes-backend/internal/middleware"
)

func Run() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", handler.Signup)
		r.Post("/login", handler.Login)
	})

	r.Route("/notes", func(r chi.Router) {
		r.Use(middleware.Auth)

		r.Post("/", handler.CreateNote)
		r.Get("/", handler.GetNotes)
		r.Get("/{id}", handler.GetNoteById)
		r.Put("/{id}", handler.UpdateNote)
		r.Delete("/{id}", handler.DeleteNote)
	})

	log.Println("Starting server on ", config.Cfg.ServerPort)
	http.ListenAndServe(":"+config.Cfg.ServerPort, r)
}
