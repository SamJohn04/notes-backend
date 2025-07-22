package app

import (
	"log"
	"net/http"

	"github.com/SamJohn04/notes-backend/internal/config"
)

func Run() {
	cfg := config.Load()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	log.Println("Starting server on ", cfg.ServerPort)
	http.ListenAndServe(":"+cfg.ServerPort, nil)
}
