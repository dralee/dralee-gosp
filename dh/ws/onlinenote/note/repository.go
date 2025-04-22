/*
仓储

2025.4.22 by dralee
*/
package note

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type NoteRepository interface {
	Save(note *Note) error
	Find(id uint32) (*Note, error)
	FindAll() ([]*Note, error)
}

type DefaultNoteRepository struct {
	db         *sql.DB
	connString string
}

func NewDefaultNoteRepository(connString string) (*DefaultNoteRepository, error) {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	repo := DefaultNoteRepository{connString: connString, db: db}
	return &repo, nil
}
