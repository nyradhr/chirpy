package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nyradhr/chirpy/internal/auth"
	"github.com/nyradhr/chirpy/internal/database"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) createChirpsHandler(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}
	type chirpRequest struct {
		Body string `json:"body"`
	}
	decoder := json.NewDecoder(r.Body)
	c := chirpRequest{}
	err = decoder.Decode(&c)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding parameters", err)
		return
	}
	if len(c.Body) > 140 {
		err = errors.New("Chirp is too long")
		respondWithError(w, http.StatusBadRequest, err.Error(), err)
		return
	}
	cleanBody := replaceBadWords(c.Body)
	params := database.CreateChirpParams{
		Body:   cleanBody,
		UserID: userID,
	}
	chirp, err := cfg.db.CreateChirp(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating chirp", err)
		return
	}
	respondWithJSON(w, http.StatusCreated, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}

func replaceBadWords(msg string) string {
	badWords := []string{"sharbert", "fornax", "kerfuffle"}
	words := strings.Split(msg, " ")
	for i, w := range words {
		if slices.Contains(badWords, strings.ToLower(w)) {
			words[i] = "****"
		}
	}
	return strings.Join(words, " ")
}
