package middleware

import (
	"errors"
	"net/http"
	"os"
	"server/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(tokenString string) (*jwt.MapClaims, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET environment variable not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, errors.New("invalid token")
}

func ExtractTokenFromCookie(c *gin.Context) string {
	if len(c.Request.CookiesNamed("auth_token")) == 0 {
		return ""
	}

	token := c.Request.CookiesNamed("auth_token")[0].Value
	return token
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := ExtractTokenFromCookie(c)
		if tokenString == "" {
			util.ErrorResponse(c.Writer, "Unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}
		_, err := ValidateJWT(tokenString)
		if err != nil {
			util.ErrorResponse(c.Writer, err.Error(), http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Next()
	}
}
