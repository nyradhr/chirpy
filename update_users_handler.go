package main

import (
	"encoding/json"
	"net/http"

	"github.com/nyradhr/chirpy/internal/auth"
	"github.com/nyradhr/chirpy/internal/database"
)

func (cfg *apiConfig) updateUsersHandler(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Error retrieving Authorization Header from Request", err)
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}
	type updateUserRequest struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	u := updateUserRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&u)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding parameters", err)
		return
	}
	hashed_password, err := auth.HashPassword(u.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error hashing password", err)
		return
	}
	params := database.UpdateUserParams{
		ID:             userID,
		Email:          u.Email,
		HashedPassword: hashed_password,
	}
	err = cfg.db.UpdateUser(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found", err)
		return
	}
	user, err := cfg.db.GetUser(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found", err)
		return
	}
	response := User{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Email:       user.Email,
		IsChirpyRed: user.IsChirpyRed,
	}
	respondWithJSON(w, http.StatusOK, response)
}
