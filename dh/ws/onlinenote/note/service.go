/*
记事服务
2015.4.22 by dralee
*/
package note

type NoteService interface {
}

type DefaultNoteService struct {
	repo NoteRepository
}
