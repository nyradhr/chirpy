package main

import (
	"net/http"
	"time"

	"github.com/nyradhr/chirpy/internal/auth"
)

func (cfg *apiConfig) getRefreshTokensHandler(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error retrieving Authorization Header from Request", err)
		return
	}
	refreshToken, err := cfg.db.GetRefreshToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Error retrieving Refresh Token from database", err)
		return
	}
	accessToken, err := auth.MakeJWT(refreshToken.UserID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating access token", err)
		return
	}
	response := struct {
		Token string `json:"token"`
	}{
		Token: accessToken,
	}
	respondWithJSON(w, http.StatusOK, response)
}
