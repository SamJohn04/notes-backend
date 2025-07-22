package app

import (
	"fmt"
	"net/http"

	"github.com/SamJohn04/notes-backend/internal/config"
)

func Run() {
	cfg := config.Load()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	fmt.Println("Starting server on ", cfg.ServerPort)
	http.ListenAndServe(":"+cfg.ServerPort, nil)
}
