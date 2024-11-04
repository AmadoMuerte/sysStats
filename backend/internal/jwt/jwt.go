package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserInfo struct {
	ID uint
}

func GenerateToken(user *UserInfo, duration time.Duration, key string, tokenType string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"iss": "FlickSynergy",
		"exp": time.Now().Add(duration).Unix(),
		"iat": time.Now().Unix(),
		"typ": tokenType,
	})

	tokenString, err := claims.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string, key string, tokenType string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if tokenType != token.Claims.(jwt.MapClaims)["typ"] {
		return nil, fmt.Errorf("invalid token type")
	}

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return token, nil
}

func ExtractUserInfo(tokenString string, secretKey []byte) (*UserInfo, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := uint(claims["sub"].(float64))
		return &UserInfo{
			ID: id,
		}, nil
	}

	return nil, fmt.Errorf("invalid token")
}
