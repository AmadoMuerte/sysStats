package jwt

import (
	"testing"
	"time"
)

const (
	testKey = "testTKey"
)

func TestGenerateToken(t *testing.T) {
	user := &UserInfo{ID: 1}

	accessToken, err := GenerateToken(user, 15*time.Minute, testKey, "access")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if accessToken == "" {
		t.Fatal("Expected a non-empty access token")
	}

	refreshToken, err := GenerateToken(user, 30*24*time.Hour, testKey, "refresh")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if refreshToken == "" {
		t.Fatal("Expected a non-empty refresh token")
	}
}

func TestVerifyToken(t *testing.T) {
	user := &UserInfo{ID: 1}
	token, err := GenerateToken(user, 15*time.Minute, testKey, "access")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Проверяем валидный токен
	_, err = VerifyToken(token, testKey)
	if err != nil {
		t.Fatalf("Expected token to be valid, got %v", err)
	}

	invalidToken := token + "invalid"
	_, err = VerifyToken(invalidToken, testKey)
	if err == nil {
		t.Fatal("Expected an error for invalid token, got none")
	}
}
