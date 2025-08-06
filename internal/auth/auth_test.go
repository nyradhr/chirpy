package auth

import (
	"testing"
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
