package middleware

import "strings"
import "github.com/gin-gonic/gin"
import "github.com/dgrijalva/jwt-go"
import "net/http"
import "fmt"

var jwtSecret = []byte("slkfjaslfjdjf!@#$!@#ASDFASDf")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		authParts := strings.SplitN(authHeader, " ", 2)
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization header"})
			return
		}

		// parse token and populate MapClaims so we can read claims easily
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(authParts[1], claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT"})
			return
		}

		if roleRaw, ok := claims["role"]; ok {
			if rs, ok := roleRaw.(string); ok {
				c.Set("role", rs)
			}
		}

		c.Next()
	}
}

func CheckAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Only Administrators can use this endpoint"})
			return
		}
		c.Next()
	}
}
