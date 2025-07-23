package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/SamJohn04/notes-backend/internal/middleware"
	"github.com/SamJohn04/notes-backend/internal/model"
	"github.com/SamJohn04/notes-backend/internal/repository"
	"github.com/go-chi/chi/v5"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note model.Note
	json.NewDecoder(r.Body).Decode(&note)
	email, err := middleware.GetUserEmail(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	note.Owner = email

	created := repository.CreateNote(note)
	json.NewEncoder(w).Encode(created)
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	email, err := middleware.GetUserEmail(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	notes := repository.GetNotesByOwner(email)
	json.NewEncoder(w).Encode(notes)
}

func GetNoteById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
	}

	email, err := middleware.GetUserEmail(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	note, err := repository.GetNoteById(id)
	if err != nil || note.Owner != email {
		http.Error(w, "Note is missing", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(note)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	var newNote model.Note

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
	}

	note, err := repository.GetNoteById(id)
	if err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&newNote)

	email, err := middleware.GetUserEmail(r)
	if err != nil || note.Owner != email || newNote.Owner != email {
		http.Error(w, "Note is missing", http.StatusNotFound)
		return
	}

	err = repository.UpdateNote(id, newNote)
	if err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
	}

	note, err := repository.GetNoteById(id)
	if err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}
	email, err := middleware.GetUserEmail(r)
	if err != nil || note.Owner != email {
		http.Error(w, "Note is missing", http.StatusNotFound)
		return
	}

	err = repository.DeleteNote(id)
	if err != nil {
		http.Error(w, "Note not found; something went wrong", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
