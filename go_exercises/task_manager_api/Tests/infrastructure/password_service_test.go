package infrastructure_test

import (
	"task_manager_api/Infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "secret123"
	hashed, err := Infrastructure.HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)
	assert.NotEqual(t, password, hashed)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "secret123"
	hashed, _ := Infrastructure.HashPassword(password)

	assert.True(t, Infrastructure.CheckPasswordHash(password, hashed))
	assert.False(t, Infrastructure.CheckPasswordHash("wrongpassword", hashed))
}
