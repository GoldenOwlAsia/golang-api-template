package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "1234"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}

	if hash == "" {
		t.Error("Hashed password is empty")
	}
}

func TestVerifyPassword(t *testing.T) {
	password := "1234"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}

	err = VerifyPassword(hash, password)
	if err != nil {
		t.Errorf("VerifyPassword returned error: %v", err)
	}

	err = VerifyPassword(hash, "passwrong")
	if err == nil {
		t.Error("VerifyPassword succeeded with incorrect password")
	}
}
