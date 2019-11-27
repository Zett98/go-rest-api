package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq" // ... should be anonymous import ???
)

// Store .. storing config here
type Store struct {
	db             *sql.DB
	UserRepository *UserRepository
}

// New ... new config
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//User ... used to manipulate users
func (s *Store) User() *UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
	}

	return s.UserRepository
}
