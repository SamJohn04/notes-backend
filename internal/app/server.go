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

	log.Println("Starting server on ", config.Cfg.ServerPort)
	http.ListenAndServe(":"+config.Cfg.ServerPort, r)
}
