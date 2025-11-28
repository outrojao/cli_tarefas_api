package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Bearer token required", http.StatusUnauthorized)
			return
		}

		validToken, err := VerifyToken(tokenString)

		if err != nil {
			http.Error(w, "verify token failed", http.StatusUnauthorized)
			return
		}
		if !validToken {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func VerifyToken(tokenString string) (bool, error) {
	secretKey := os.Getenv("JWT_SECRET")
	subHeader := os.Getenv("JWT_HEADER_SUB")
	issHeader := os.Getenv("JWT_HEADER_ISS")

	if secretKey == "" || subHeader == "" || issHeader == "" {
		return false, nil
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if sub, ok := claims["sub"].(string); !ok || sub != subHeader {
			return false, nil
		}

		if iss, ok := claims["iss"].(string); !ok || iss != issHeader {
			return false, nil
		}

		if iat, ok := claims["iat"].(float64); !ok || int64(iat) <= 0 {
			return false, nil
		}

		return true, nil
	}

	return false, nil
}

func CreateToken() (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	sub := os.Getenv("JWT_HEADER_SUB")
	iss := os.Getenv("JWT_HEADER_ISS")

	if secretKey == nil || sub == "" || iss == "" {
		return "", fmt.Errorf("invalid jwt header")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"iss": iss,
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}
