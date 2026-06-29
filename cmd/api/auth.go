package main

import "net/http"

func (a *app) signUpHandler(w http.ResponseWriter, r *http.Request) {
	token, err := a.authenticator.GenerateToken(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, "could not create session")
		return
	}

	respondJSON(w, http.StatusCreated, map[string]string{"token": token})
}

func (a *app) logInHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		respondError(w, http.StatusBadRequest, "token query parameter is required")
		return
	}

	err := a.authenticator.ValidateToken(r.Context(), token)
	if err != nil {
		respondError(w, http.StatusUnauthorized, "invalid or expired token")
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"status": "authenticated"})
}
