package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/nyradhr/chirpy/internal/database"
)

func (cfg *apiConfig) listChirpsHandler(w http.ResponseWriter, r *http.Request) {
	var dbChirps []database.Chirp
	var err error
	authorIDString := r.URL.Query().Get("author_id")
	if authorIDString != "" {
		authorID, err := uuid.Parse(authorIDString)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author ID format", err)
			return
		}
		dbChirps, err = cfg.db.ListChirpsByAuthor(r.Context(), authorID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error listing chirps", err)
			return
		}
	} else {
		dbChirps, err = cfg.db.ListChirps(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error listing chirps", err)
			return
		}
	}
	chirps := make([]Chirp, len(dbChirps))
	for i, c := range dbChirps {
		chirps[i] = Chirp{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			Body:      c.Body,
			UserID:    c.UserID,
		}
	}
	respondWithJSON(w, http.StatusOK, chirps)
}
