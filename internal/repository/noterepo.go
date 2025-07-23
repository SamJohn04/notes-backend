package repository

import (
	"errors"
	"github.com/SamJohn04/notes-backend/internal/model"
)

var notes = map[int]model.Note{} // title -> note TODO database
var nextId = 1                   // TODO database

func CreateNote(note model.Note) model.Note {
	note.Id = nextId
	nextId++
	notes[note.Id] = note
	return note
}

func GetNotesByOwner(email string) []model.Note {
	ownerNotes := []model.Note{}
	for _, note := range notes {
		if note.Owner == email {
			ownerNotes = append(ownerNotes, note)
		}
	}
	return ownerNotes
}

func GetNoteById(id int) (model.Note, error) {
	note, ok := notes[id]
	if !ok {
		return model.Note{}, errors.New("note not found")
	}
	return note, nil
}

func UpdateNote(id int, updated model.Note) error {
	if _, exists := notes[id]; !exists {
		return errors.New("note not found")
	}
	updated.Id = id
	notes[id] = updated
	return nil
}

func DeleteNote(id int) error {
	if _, exists := notes[id]; !exists {
		return errors.New("note not found")
	}
	delete(notes, id)
	return nil
}
