package main

import (
	"backend-go/application"
	"net/http"
)

// @Tags Продукты
// @Summary Создать продукт
// @Description Создать продукт
// @Accept json
// @Produce json
// @Param JSON body api.CreateProductRequestDto true "CreateProductRequestDto"
// @Success 200 {object} string
// @Router /products [POST]
func CreateProductHanlder(
	w http.ResponseWriter,
	r *http.Request) {

	err := application.CreateProductHanlder(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
