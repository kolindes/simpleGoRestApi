package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kolindes/simpleRestApi/internal/error"
	"github.com/kolindes/simpleRestApi/internal/models"
)

func PathNotFound(w http.ResponseWriter, r *http.Request) {
	response := models.NewResponse()
	response.Message = error.AccessDenied

	accessToken := r.Header.Get("Authorization")
	if accessToken == "" {
		response.Error = error.Unauthorized
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Message = error.NotFound
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}
