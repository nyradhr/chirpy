package main

import "net/http"

func (cfg *apiConfig) listChirpsHandler(w http.ResponseWriter, r *http.Request) {
	result, err := cfg.db.ListChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error listing chirps", err)
		return
	}
	chirps := make([]Chirp, len(result))
	for i, c := range result {
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
