package Infrastructure

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("slkfjaslfjdjf!@#$!@#ASDFASDf")

func ValidateToken(tokenString string, claims jwt.Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
}

func GenerateJWT(username string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["role"] = role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
