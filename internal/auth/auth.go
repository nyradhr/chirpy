package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    "chirpy",
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
			Subject:   userID.String(),
		},
	)
	token, err := t.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	claims := &jwt.RegisteredClaims{}
	t, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}
	id, err := t.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}
	userID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, err
	}
	return userID, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	auth := headers.Get("Authorization")
	if auth == "" {
		return "", errors.New("authorization header not found")
	}
	f := func(c rune) bool {
		return c == ' '
	}
	parts := strings.FieldsFunc(auth, f)
	if len(parts) != 2 {
		return "", errors.New("malformed authorization header")
	}
	if parts[0] != "Bearer" {
		return "", errors.New("wrong authorization header prefix")
	}
	token := parts[1]
	return token, nil
}

func MakeRefreshToken() (string, error) {
	random := make([]byte, 32)
	_, err := rand.Read(random)
	if err != nil {
		return "", err
	}
	refreshToken := hex.EncodeToString(random)
	return refreshToken, nil
}
