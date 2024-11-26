package handlers

import (
	"encoding/json"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/api/requests"
	"net/http"
	"strings"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewAuth(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usersDB := MongoDB(r).Users()
	user, err := usersDB.FindByEmail(req.Email)
	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "User not found",
			})
			return
		}
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	if user.PasswordHash != req.PasswordHash {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid credentials",
		})
		return
	}

	// Якщо автентифікація успішна
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User authenticated successfully",
		"userID":  user.ID.Hex(),
	})
}
