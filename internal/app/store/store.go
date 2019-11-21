package store

import (
	"database/sql"
	_ "github.com/lib/pq" // ... should be anonymous import ???
)

// Store .. storing config here
type Store struct {
	config         *Config
	db             *sql.DB
	UserRepository *UserRepository
}

// New ... new config
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ... opens db connection
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

//Close ... closes db conn
func (s *Store) Close() {
	s.db.Close()
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
