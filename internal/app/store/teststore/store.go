package teststore

import (
	"go-rest-api/internal/app/model"
	"go-rest-api/internal/app/store"
)

// Store .. storing config here
type Store struct {
	UserRepository *UserRepository
}

// New ... new config
func New() *Store {
	return &Store{}
}

//User ... used to manipulate users
func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.UserRepository
}
