package handlers

import (
	"backend-go/application/createProduct"

	"backend-go/infrastructure/persistence"

	"net/http"
)

// @Tags Продукты
// @Summary Создать продукт
// @Description Создать продукт
// @Accept json
// @Produce json
// @Param JSON body createProduct.CreateProductRequestDto true "CreateProductRequestDto"
// @Success 200 {object} string
// @Router /products [POST]
func CreateProductHanlder(
	w http.ResponseWriter,
	r *http.Request) {

	repo, err := persistence.NewRepositories()

	if err != nil {
		panic(err)
	}

	defer repo.Close()

	err = createProduct.CreateProductHanlder(w, r, repo.Product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
