package infrastructure_test

import (
	"task_manager_api/Infrastructure"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	token, err := Infrastructure.GenerateJWT("testuser", "admin")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateToken(t *testing.T) {
	tokenString, _ := Infrastructure.GenerateJWT("testuser", "admin")
	claims := jwt.MapClaims{}
	token, err := Infrastructure.ValidateToken(tokenString, claims)

	assert.NoError(t, err)
	assert.True(t, token.Valid)
	assert.Equal(t, "testuser", claims["username"])
	assert.Equal(t, "admin", claims["role"])
}

func TestValidateInvalidToken(t *testing.T) {
	claims := jwt.MapClaims{}
	_, err := Infrastructure.ValidateToken("invalid.token.string", claims)
	assert.Error(t, err)
}
