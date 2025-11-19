package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

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
