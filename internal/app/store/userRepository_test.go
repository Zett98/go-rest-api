package store

import "testing"

import "go-rest-api/internal/app/model"

import "github.com/stretchr/testify/assert"

// TestUserRepositoryCreate ... tests creation of user
func TestUserRepositoryCreate(t *testing.T) {
	s, teardown := TestStore(t, databaseURL)

	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@zalupa.mail",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

// TestUserRepositoryFindByEmail ... tests find user by id
func TestUserRepositoryFindByEmail(t *testing.T) {
	s, teardown := TestStore(t, databaseURL)

	defer teardown("users")

	email := "user@zalupa.mail"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "user@zalupa.mail",
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
