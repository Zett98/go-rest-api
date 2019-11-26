package model

import "testing"

// TestUser ... creates new User
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "adminadmin",
	}
}
