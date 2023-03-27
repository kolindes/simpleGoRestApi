package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kolindes/simpleRestApi/internal/database"
	"github.com/kolindes/simpleRestApi/internal/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response

	response.Message = "Not registered"

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = "Bad request: " + err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	if user.Username == "" || user.Password == "" || user.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = "Username, password and Email are required."
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := user.SetPassword(user.Password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Error = "Internal error: " + err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	err := database.AddUser(user.Username, user.Email, user.HashedPassword)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Error = "The user was not added to the database, error: " + err.Error()
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
	w.WriteHeader(http.StatusCreated)
	response.Message = "Registered"
	response.Data["access_token"] = token
	json.NewEncoder(w).Encode(response)
}
