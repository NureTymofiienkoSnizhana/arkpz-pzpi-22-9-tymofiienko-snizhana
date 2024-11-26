package handlers

import (
	"encoding/json"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/api/requests"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewRegistration(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := data.User{
		ID:           primitive.NewObjectID(),
		FullName:     req.FullName,
		Email:        req.Email,
		Role:         "user",
		PasswordHash: req.PasswordHash,
	}

	usersDB := MongoDB(r).Users()

	err = usersDB.Insert(&user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key error collection") {
			w.WriteHeader(http.StatusConflict)
			return
		}
		http.Error(w, "Failed to save user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
		"userID":  user.ID.Hex(),
	})
}
