package store

import "go-rest-api/internal/app/model"

// userRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
