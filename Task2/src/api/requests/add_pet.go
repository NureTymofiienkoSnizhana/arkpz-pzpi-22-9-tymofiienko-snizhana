package requests

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type AddPet struct {
	Name    string `json:"name"`
	Speices string `json:"speices"`
	Breed   string `json:"breed"`
	Age     int    `json:"age"`
}

func NewPet(r *http.Request) (*AddPet, error) {
	bodyReader := r.Body
	if bodyReader == nil {
		return nil, errors.New("missing body")
	}

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, err
	}

	var pet AddPet
	err = json.Unmarshal(body, &pet)
	if err != nil {
		return nil, err
	}

	return &pet, nil
}
