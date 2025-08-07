package main

import (
	"net/http"

	"github.com/nyradhr/chirpy/internal/auth"
)

func (cfg *apiConfig) revokeRefreshTokensHandler(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error retrieving Authorization Header from Request", err)
		return
	}
	err = cfg.db.RevokeRefreshToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error revoking refresh token", err)
		return
	}
	respondWithJSON(w, http.StatusNoContent, nil)
}
