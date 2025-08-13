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
	userId, err := middleware.GetUserId(r)
	if err != nil {
		log.Println("Unauthorized in CreateNote:", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	note.Owner = userId

	err = repository.CreateNote(note)
	if err != nil {
		log.Println("DB error in CreateNote:", err)
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	userId, err := middleware.GetUserId(r)
	if err != nil {
		log.Println("Unauthorized in GetNotes:", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	notes, err := repository.GetNotesByOwner(userId)
	if err != nil {
		log.Println("DB error in GetNotes:", err)
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(notes)
}

func GetNoteById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("ID is missing in GetNoteById:", err)
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
	}

	userId, err := middleware.GetUserId(r)
	if err != nil {
		log.Println("Unauthorized in GetNoteById:", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	note, err := repository.GetNoteById(id)
	if err != nil || userId != note.Owner {
		log.Println("Note is missing in GetNoteById:", err)
		http.Error(w, "Note is missing", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(note)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	var newNote model.Note

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("ID is missing in UpdateNote:", err)
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
	}

	json.NewDecoder(r.Body).Decode(&newNote)

	userId, err := middleware.GetUserId(r)
	if err != nil {
		log.Println("Invalid user id in UpdateNote:", err)
		http.Error(w, "Invalid user id", http.StatusUnauthorized)
		return
	}

	newNote.Owner = userId
	err = repository.UpdateNote(id, newNote)
	if err != nil {
		log.Println("Note not found in UpdateNote:", err)
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("ID missing in DeleteNote:", err)
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
	}

	userId, err := middleware.GetUserId(r)
	if err != nil {
		log.Println("Unauthorized in DeleteNote:", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = repository.DeleteNote(id, userId)
	if err != nil {
		log.Println("Note not found in DeleteNote:", err)
		http.Error(w, "Note not found; something went wrong", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
