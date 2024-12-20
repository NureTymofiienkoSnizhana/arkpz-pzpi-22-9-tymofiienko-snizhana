package handlers

import (
	"encoding/json"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/api/requests"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func UpdateDevice(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewUpdateDevice(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateFields := bson.M{
		"pet_id":         req.PetID,
		"status":         req.Status,
		"last_sync_time": req.LastSyncTime,
	}

	if req.ID.IsZero() {
		http.Error(w, "Invalid device ID", http.StatusBadRequest)
		return
	}

	devicesDB := MongoDB(r).Devices()

	err = devicesDB.Update(req.ID, updateFields)
	if err != nil {
		http.Error(w, "Failed to update device", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Device updated successfully",
		"deviceID": req.ID.Hex(),
	})
}
