package note

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	Id     int
	Value  string
	Author string
}

type NoteRepository interface {
	Health() (string, error)
	GetAllNotes() ([]Note, error)
	AddNote(Note) (Note, error)
	DeleteNote(id int) error
}

type sqliteNoteRepository struct {
	db *sql.DB
}

func NewNoteRepository(db *sql.DB) NoteRepository {
	// attempt to init schema
	data, err := os.ReadFile("internal/repositories/note/schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(data))
	if err != nil {
		log.Fatal(err)
	}

	s := &sqliteNoteRepository{db: db}
	return s
}

func (s *sqliteNoteRepository) Health() (string, error) {
	cx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	_, err := s.db.QueryContext(cx, "select 1;")
	if err != nil {
		return "", err
	}

	return "note repository is healthy", nil
}

func (s *sqliteNoteRepository) GetAllNotes() ([]Note, error) {
	cx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	results, err := s.db.QueryContext(cx, "select id, value, author from notes;")
	if err != nil {
		return []Note{}, err
	}

	notes := []Note{}

	for results.Next() {
		n := Note{}
		err := results.Scan(&n.Id, &n.Value, &n.Author)
		if err != nil {
			continue
		}
		notes = append(notes, n)
	}

	return notes, nil
}

func (s *sqliteNoteRepository) AddNote(note Note) (Note, error) {
	cx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stmt, err := s.db.PrepareContext(cx, "insert into notes (value, author) values (?, ?);")
	if err != nil {
		return note, err
	}

	result, err := stmt.ExecContext(cx, note.Value, note.Author)
	if err != nil {
		return note, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return note, err
	}
	if rowsAffected != 1 {
		return note, fmt.Errorf("unable to create note")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return note, err
	}
	note.Id = int(id)

	return note, nil
}

func (s *sqliteNoteRepository) DeleteNote(id int) error {
	cx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stmt, err := s.db.PrepareContext(cx, "delete from notes where id = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(cx, id)
	if err != nil {
		return err
	}

	return nil
}
