package application

import (
	"backend-go/api"
	"backend-go/infrastructure"
	"encoding/json"
	"net/http"
)

func CreateProductHanlder(
	w http.ResponseWriter,
	r *http.Request) {

	var requestDto api.CreateProductDto

	err := json.NewDecoder(r.Body).Decode(&requestDto)

	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(requestDto)

	infrastructure.Create("", "", 0)
}
