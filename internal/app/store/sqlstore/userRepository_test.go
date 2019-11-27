package sqlstore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go-rest-api/internal/app/model"
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
	u1 := model.TestUser(t)
	_, err := s.User().FindByEmail(u1.Email)

	s.User().Create(u1)
	u2, err := s.User().FindByEmail(u1.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)

}
