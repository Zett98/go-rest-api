package teststore

import (
	"github.com/stretchr/testify/assert"
	"go-rest-api/internal/app/model"
	"go-rest-api/internal/app/store"
	"testing"
)

// TestUserRepositoryCreate ... tests creation of user
func TestUserRepositoryCreate(t *testing.T) {
	s := New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

// TestUserRepositoryFindByEmail ... tests find user by id
func TestUserRepositoryFindByEmail(t *testing.T) {
	s := New()
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}

// TestUserRepositoryFind ... tests find user by id
func TestUserRepositoryFind(t *testing.T) {

	s := New()
	u1 := model.TestUser(t)

	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)

}
