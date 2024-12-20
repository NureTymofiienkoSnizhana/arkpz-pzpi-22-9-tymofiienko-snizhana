package handlers

import (
	"encoding/json"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/api/requests"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewUpdateUser(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateFields := bson.M{
		"full_name":     req.FullName,
		"email":         req.Email,
		"role":          req.Role,
		"password_hash": req.PasswordHash,
		"pets_id":       req.PetsID,
	}

	if req.ID.IsZero() {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	usersDB := MongoDB(r).Users()

	err = usersDB.Update(req.ID, updateFields)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User updated successfully",
		"userID":  req.ID.Hex(),
	})
}
