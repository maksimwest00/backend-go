package application

import (
	"backend-go/api"
	"backend-go/infrastructure"
	"encoding/json"
	"fmt"
	"net/http"
)

// @Tags Продукты
// @Summary Создать продукт
// @Description Создать продукт
// @Accept json
// @Produce json
// @Param JSON body api.CreateProductRequestDto true "CreateProductDto"
// @Success 200 {object} string
// @Router /products [POST]
func CreateProductHanlder(
	w http.ResponseWriter,
	r *http.Request) {

	var requestDto api.CreateProductRequestDto

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&requestDto)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err = requestDto.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := infrastructure.Create(requestDto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Продукт успешно создан id: %d", *res)
}
