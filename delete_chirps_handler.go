package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/nyradhr/chirpy/internal/auth"
)

func (cfg *apiConfig) deleteChirpsHandler(w http.ResponseWriter, r *http.Request) {
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
	chirpID := r.PathValue("chirpID")
	id, err := uuid.Parse(chirpID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID format", err)
		return
	}
	c, err := cfg.db.GetChirp(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}
	chirp := Chirp{
		ID:     c.ID,
		UserID: c.UserID,
	}
	if userID != chirp.UserID {
		respondWithError(w, http.StatusForbidden, "Cannot delete other users' chirps", err)
		return
	}
	err = cfg.db.DeleteChirp(r.Context(), chirp.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error deleting chirp", err)
	}
	respondWithJSON(w, http.StatusNoContent, nil)
}
