package repository

import (
	"errors"

	"github.com/SamJohn04/notes-backend/internal/config"
	"github.com/SamJohn04/notes-backend/internal/model"
)

func CreateNote(note model.Note) error {
	_, err := config.DB.Exec(
		"INSERT INTO notes (user_id, title, body) VALUES (?, ?, ?)",
		note.Owner, note.Title, note.Body,
	)
	return err
}

func GetNotesByOwner(userId int) ([]model.Note, error) {
	ownerNotes := []model.Note{}

	rows, err := config.DB.Query(
		"SELECT id, title, body FROM notes where user_id=?",
		userId,
	)
	if err != nil {
		return ownerNotes, errors.New("DB error")
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title, body string

		if err = rows.Scan(&id, &title, &body); err != nil {
			return ownerNotes, errors.New("DB error")
		}

		ownerNotes = append(ownerNotes, model.Note{
			Id:    id,
			Title: title,
			Body:  body,
			Owner: userId,
		})
	}

	return ownerNotes, nil
}

func GetNoteById(id int) (model.Note, error) {
	var note model.Note
	note.Id = id

	err := config.DB.QueryRow(
		"SELECT user_id, title, body FROM notes WHERE id=?",
		id,
	).Scan(note.Owner, note.Title, note.Body)

	return note, err
}

func UpdateNote(id int, updated model.Note) error {
	res, err := config.DB.Exec(
		"UPDATE notes SET title = ?, body = ? WHERE id = ? AND user_id = ?",
		updated.Title, updated.Body, id, updated.Owner,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New("note not found or unauthorized")
	}

	return nil
}

func DeleteNote(noteId, userId int) error {
	res, err := config.DB.Exec(
		"DELETE FROM notes WHERE id = ? AND user_id = ?",
		noteId, userId,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New("note not found or unauthorized")
	}

	return nil
}
