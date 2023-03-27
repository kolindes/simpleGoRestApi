package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kolindes/simpleRestApi/internal/database"
	"github.com/kolindes/simpleRestApi/internal/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("LOGIN: LoginHandler, credentials:", credentials)

	user, err := database.GetUserByUsername(credentials.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Println("LOGIN: LoginHandler, user(db):", user)
	if !user.CheckPassword(credentials.Password) {
		http.Error(w, "Invalid credentials:"+string(user.Password)+"\t"+string(credentials.Password), http.StatusUnauthorized)
		return
	}

	token, err := generateToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"message": "Authorized", "access_token": token})
}
