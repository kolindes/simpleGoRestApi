package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kolindes/simpleRestApi/internal/database"
	"github.com/kolindes/simpleRestApi/internal/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := models.NewResponse()

	response.Message = "Unauthorized"

	var credentials models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	if credentials.Username == "" || credentials.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = "Invalid username or password."
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := database.GetUserByUsername(credentials.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		response.Error = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	if !user.CheckPassword(credentials.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		response.Error = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	token, err := generateToken(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Error = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)
	response.Message = "Authorized"
	response.Data["access_token"] = token
	json.NewEncoder(w).Encode(response)
}
