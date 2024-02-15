package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fadur/hello/internal/models"
)

// list all users

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	users, err := models.ListUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload := map[string]interface{}{
		"users": users,
	}
	json.NewEncoder(w).Encode(payload)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = user.Save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload := map[string]interface{}{
		"user": user,
	}
	json.NewEncoder(w).Encode(payload)
}
