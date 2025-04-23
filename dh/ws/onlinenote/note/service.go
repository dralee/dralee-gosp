/*
记事服务
2015.4.22 by dralee
*/
package note

type NoteService interface {
	Save(note *Note) error
	Find(id uint32) (*Note, error)
	FindByUserId(userId uint32) ([]*Note, error)
	FindAll() ([]*Note, error)
	Update(note *Note) error
	Delete(id uint32) error
}

type DefaultNoteService struct {
	repo NoteRepository
}

func NewDefaultNoteService(repo NoteRepository) *DefaultNoteService {
	return &DefaultNoteService{repo: repo}
}

func (r *DefaultNoteService) Save(note *Note) error {
	return r.repo.Save(note)
}

func (r *DefaultNoteService) Find(id uint32) (*Note, error) {
	return r.repo.Find(id)
}

func (r *DefaultNoteService) FindByUserId(userId uint32) ([]*Note, error) {
	return r.repo.FindByUserId(userId)
}

func (r *DefaultNoteService) FindAll() ([]*Note, error) {
	return r.repo.FindAll()
}

func (r *DefaultNoteService) Update(note *Note) error {
	return r.repo.Update(note)
}

func (r *DefaultNoteService) Delete(id uint32) error {
	return r.repo.Delete(id)
}
