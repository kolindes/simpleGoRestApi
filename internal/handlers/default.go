package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kolindes/simpleRestApi/internal/models"
)

func PathNotFound(w http.ResponseWriter, _ *http.Request) {
	response := models.NewResponse()
	response.Message = "Not found"
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}
