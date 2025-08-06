package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestHashPassword(t *testing.T) {
	password := "testpassword123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if hash == "" {
		t.Fatal("Expected hash to not be empty")
	}
	if hash == password {
		t.Fatal("Hash should not equal original password")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "testpassword123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}
	err = CheckPasswordHash(password, hash)
	if err != nil {
		t.Fatalf("Expected password to match hash, got error: %v", err)
	}
	err = CheckPasswordHash("wrongpassword", hash)
	if err == nil {
		t.Fatal("Expected error for wrong password, got nil")
	}
}

func TestJWT(t *testing.T) {
	userId := uuid.New()
	tokenSecret := "test-secret"
	token, err := MakeJWT(userId, tokenSecret, time.Minute)
	if err != nil {
		t.Fatalf("Failed to create jwt: %v", err)
	}
	id, err := ValidateJWT(token, tokenSecret)
	if err != nil {
		t.Fatalf("Failed to validate jwt: %v", err)
	}
	if id != userId {
		t.Fatal("UserID from validation does not match with original")
	}
}

func TestJWTExpired(t *testing.T) {
	userId := uuid.New()
	tokenSecret := "test-secret"
	token, err := MakeJWT(userId, tokenSecret, -time.Second)
	if err != nil {
		t.Fatalf("Failed to create jwt: %v", err)
	}
	_, err = ValidateJWT(token, tokenSecret)
	if err == nil {
		t.Fatal("Expected error for expired token, got nil")
	}
}

func TestJWTWrongSecret(t *testing.T) {
	userId := uuid.New()
	tokenSecret := "test-secret"
	token, err := MakeJWT(userId, tokenSecret, time.Minute)
	if err != nil {
		t.Fatalf("Failed to create jwt: %v", err)
	}
	wrongSecret := "wrong-secret"
	_, err = ValidateJWT(token, wrongSecret)
	if err == nil {
		t.Fatal("Expected error for wrong secret, got nil")
	}
}
