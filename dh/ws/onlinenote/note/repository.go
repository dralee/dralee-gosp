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
	FindByUserId(userId uint32) ([]*Note, error)
	FindAll() ([]*Note, error)
	Update(note *Note) error
	Delete(id uint32) error
	Close()
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

func (r *DefaultNoteRepository) Save(note *Note) error {
	var sql = "INSERT INTO `note` (`name`, `content`, `user_id`, `creation_time`, `last_modification_time`) VALUES (?, ?, ?, ?, ?);"
	_, err := r.db.Exec(sql, note.Name, note.Content, note.CreatorId, note.CreationTime, note.LastModificationTime)
	return err
}

func (r *DefaultNoteRepository) Find(id uint32) (*Note, error) {
	var sql = "SELECT `id`, `name`, `content`, `user_id`, `creation_time`, `last_modification_time` FROM `note` WHERE `id` = ?;"
	var note Note
	err := r.db.QueryRow(sql, id).Scan(&note.Id, &note.Name, &note.Content, &note.CreatorId, &note.CreationTime, &note.LastModificationTime)
	return &note, err
}

func (r *DefaultNoteRepository) FindByUserId(userId uint32) ([]*Note, error) {
	var sql = "SELECT `id`, `name`, `content`, `user_id`, `creation_time`, `last_modification_time` FROM `note` WHERE `user_id` = ?;"
	rows, err := r.db.Query(sql, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var notes []*Note
	for rows.Next() {
		var note Note
		err := rows.Scan(&note.Id, &note.Name, &note.Content, &note.CreatorId, &note.CreationTime, &note.LastModificationTime)
		if err != nil {
			return nil, err
		}
		notes = append(notes, &note)
	}
	return notes, nil
}

func (r *DefaultNoteRepository) FindAll() ([]*Note, error) {
	var sql = "SELECT `id`, `name`, `content`, `user_id`, `creation_time`, `last_modification_time` FROM `note`;"
	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var notes []*Note
	for rows.Next() {
		var note Note
		err := rows.Scan(&note.Id, &note.Name, &note.Content, &note.CreatorId, &note.CreationTime, &note.LastModificationTime)
		if err != nil {
			return nil, err
		}
		notes = append(notes, &note)
	}
	return notes, nil
}

func (r *DefaultNoteRepository) Update(note *Note) error {
	var sql = "UPDATE `note` SET `name` = ?, `content` = ?, `last_modification_time` = ? WHERE `id` = ?;"
	_, err := r.db.Exec(sql, note.Name, note.Content, note.LastModificationTime, note.Id)
	return err
}

func (r *DefaultNoteRepository) Delete(id uint32) error {
	var sql = "DELETE FROM `note` WHERE `id` = ?;"
	_, err := r.db.Exec(sql, id)
	return err
}

func (r *DefaultNoteRepository) Close() {
	r.db.Close()
}
