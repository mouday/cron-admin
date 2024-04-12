package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(userId string, secretString string) string {
	// 秘钥
	secret := []byte(secretString)

	// 参数
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
		})

	tokenString, _ := token.SignedString(secret)

	return tokenString
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZ2UiOjIwLCJuYW1lIjoiVG9tIn0.bU-8W6kUQM55ZT-mviisDnGja2nOmvYXkyWirrFMUf0
}

func VerifyToken(tokenString string, secretString string) bool {
	_, err := ParseToken(tokenString, secretString)

	if err != nil {
		return false
	} else {
		return true
	}
}

func ParseToken(tokenString string, secretString string) (string, error) {
	// 秘钥
	secret := []byte(secretString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["userId"].(string), nil
	} else {
		return "", nil
	}
}
