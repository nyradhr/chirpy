package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nyradhr/chirpy/internal/auth"
)

func (cfg *apiConfig) loginHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password         string `json:"password"`
		Email            string `json:"email"`
		ExpiresInSeconds int    `json:"expires_in_seconds"`
	}
	decoder := json.NewDecoder(r.Body)
	l := parameters{}
	err := decoder.Decode(&l)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding parameters", err)
		return
	}
	u, err := cfg.db.GetUser(r.Context(), l.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}
	user := User{
		ID:             u.ID,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
	}
	err = auth.CheckPasswordHash(l.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}
	expirationTime := time.Hour
	if l.ExpiresInSeconds > 0 {
		requestedDuration := time.Duration(l.ExpiresInSeconds) * time.Second
		expirationTime = min(requestedDuration, time.Hour)
	}
	token, err := auth.MakeJWT(user.ID, cfg.jwtSecret, expirationTime)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating jwt token", err)
		return
	}
	type UserResponse struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email     string    `json:"email"`
		Token     string    `json:"token"`
	}
	respondWithJSON(w, http.StatusOK, UserResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		Token:     token,
	})
}
