package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kolindes/simpleRestApi/internal/models"
	"github.com/kolindes/simpleRestApi/internal/svcerr"
)

func PathNotFound(w http.ResponseWriter, r *http.Request) {
	response := models.NewResponse()
	response.Message = svcerr.AccessDenied

	accessToken := r.Header.Get("Authorization")
	if accessToken == "" {
		response.Error = svcerr.Unauthorized
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Message = svcerr.NotFound
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}
