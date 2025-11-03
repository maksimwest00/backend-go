package handlers

import (
	"backend-go/infrastructure/persistence"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Tags Продукты
// @Summary Получить продукт по id
// @Description Получить продукт по id
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} domain.Product
// @Router /products/{id} [GET]
func GetProductByIdHandler(
	w http.ResponseWriter,
	r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	repo, err := persistence.NewRepositories()

	if err != nil {
		panic(err)
	}

	defer repo.Close()

	number, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	product, err := repo.Product.GetProductById(&number)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
