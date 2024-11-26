package handlers

import (
	"encoding/json"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/api/requests"
	"net/http"
)

func PetInfo(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewPetID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.ID.IsZero() {
		http.Error(w, "Invalid pet ID", http.StatusBadRequest)
		return
	}

	petsDB := MongoDB(r).Pets()

	pet, err := petsDB.Get(req.ID)
	if err != nil {
		http.Error(w, "Failed to retrieve pet", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(pet)
}
