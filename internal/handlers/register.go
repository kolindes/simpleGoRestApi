package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kolindes/simpleRestApi/internal/database"
	"github.com/kolindes/simpleRestApi/internal/models"
	"github.com/kolindes/simpleRestApi/internal/svcerr"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := models.NewResponse()

	response.Message = svcerr.NotRegistered

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = svcerr.BadRequest
		json.NewEncoder(w).Encode(response)
		return
	}

	if user.Username == "" || user.Password == "" || user.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = svcerr.InvalidRegisterCredentials
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := user.SetPassword(user.Password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Error = svcerr.InternalServerError
		json.NewEncoder(w).Encode(response)
		return
	}

	// gorm database add user
	err := database.AddUser(user.Username, user.Email, user.HashedPassword)

	if err != nil {
		if err.Error() == svcerr.UserAlreadyExists {
			w.WriteHeader(http.StatusConflict)
			response.Error = svcerr.UserAlreadyExists
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		response.Error = svcerr.NotRegistered + ". Unexpected error:<" + err.Error() + ">"
		json.NewEncoder(w).Encode(response)
		return
	}

	token, err := generateToken(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Error = svcerr.InternalServerError
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusCreated)
	response.Message = "registered"
	response.Data["access_token"] = token
	json.NewEncoder(w).Encode(response)
}
