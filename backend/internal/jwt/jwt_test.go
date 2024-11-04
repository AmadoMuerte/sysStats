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
	t.Log("Access token generated successfully")

	refreshToken, err := GenerateToken(user, 30*24*time.Hour, testKey, "refresh")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if refreshToken == "" {
		t.Fatal("Expected a non-empty refresh token")
	}
	t.Log("Refresh token generated successfully")
}

func TestVerifyToken(t *testing.T) {
	user := &UserInfo{ID: 1}
	token, err := GenerateToken(user, 15*time.Minute, testKey, "access")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = VerifyToken(token, testKey, "access")
	if err != nil {
		t.Fatalf("Expected token to be valid, got %v", err)
	}
	t.Log("Token verified successfully")

	invalidToken := token + "invalid"
	_, err = VerifyToken(invalidToken, testKey, "access")
	if err == nil {
		t.Fatal("Expected an error for invalid token, got none")
	}
	t.Log("Invalid token correctly identified")
}

func TestExtractUserInfo(t *testing.T) {
	user := &UserInfo{ID: 1}
	token, err := GenerateToken(user, 15*time.Minute, testKey, "access")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = ExtractUserInfo(token, []byte(testKey))
	if err != nil {
		t.Fatalf("Expected token to be valid, got %v", err)
	}
	t.Log("UserInfo extracted successfully")

	invalidToken := token + "invalid"
	_, err = ExtractUserInfo(invalidToken, []byte(testKey))
	if err == nil {
		t.Fatal("Expected an error for invalid token, got none")
	}
	t.Log("Invalid token correctly identified while extracting user info")
}
