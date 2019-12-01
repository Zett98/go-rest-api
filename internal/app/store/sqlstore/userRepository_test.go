package sqlstore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go-rest-api/internal/app/model"
	"go-rest-api/internal/app/store"
)

// TestUserRepositoryCreate ... tests creation of user
func TestUserRepositoryCreate(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)

	defer teardown("users")

	s := New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

// TestUserRepositoryFindByEmail ... tests find user by id
func TestUserRepositoryFindByEmail(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)

	defer teardown("users")

	s := New(db)
	u := model.TestUser(t)
	_, err := s.User().FindByEmail(u.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u)
	u, err = s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
