package handlers

import (
	"encoding/json"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := map[string]interface{}{
		"status": "OK",
	}
	json.NewEncoder(w).Encode(data)
}
