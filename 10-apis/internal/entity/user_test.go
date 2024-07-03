package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "sJw0H@example.com", "password123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "sJw0H@example.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "sJw0H@example.com", "password123")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("password123"))
	assert.False(t, user.ValidatePassword("password456"))
	assert.NotEqual(t, "password123", user.Password)
}
