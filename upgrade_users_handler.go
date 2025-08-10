package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/nyradhr/chirpy/internal/auth"
)

func (cfg *apiConfig) upgradeUsersHandler(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error(), err)
		return
	}
	if apiKey != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "Api key does not match", err)
		return
	}
	type upgradeUserRequest struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}
	u := upgradeUserRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&u)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding parameters", err)
		return
	}
	if u.Event != "user.upgraded" {
		respondWithJSON(w, http.StatusNoContent, nil)
		return
	}
	userID, err := uuid.Parse(u.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid UserID format", err)
		return
	}
	err = cfg.db.UpgradeUser(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found", err)
		return
	}
	respondWithJSON(w, http.StatusNoContent, nil)
}
