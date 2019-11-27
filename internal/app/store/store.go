package store

import "go-rest-api/internal/app/store/sqlstore"

// Store ...
type Store interface {
	User() sqlstore.UserRepository
}
