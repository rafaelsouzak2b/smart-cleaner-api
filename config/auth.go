package config

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = map[string][]byte{
	"cleaner": []byte(Environment.CleanerToken),
}

type Claims struct {
	Id   string `json:"cleaner_id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(id string, role string) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &Claims{
		Id:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey[role])
}

func JWTMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey[role], nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("id", claims.Id)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// tokenString := c.GetHeader("Authorization")
		// if tokenString == "" {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		// 	c.Abort()
		// 	return
		// }
		// if tokenString != Environment.DefaultToken {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		// 	c.Abort()
		// 	return
		// }
		c.Next()
	}
}
