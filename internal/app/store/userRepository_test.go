package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go-rest-api/internal/app/model"
)

// TestUserRepositoryCreate ... tests creation of user
func TestUserRepositoryCreate(t *testing.T) {
	s, teardown := TestStore(t, databaseURL)

	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

// TestUserRepositoryFindByEmail ... tests find user by id
func TestUserRepositoryFindByEmail(t *testing.T) {
	s, teardown := TestStore(t, databaseURL)

	defer teardown("users")

	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
