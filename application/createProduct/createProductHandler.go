package createProduct

import (
	repository "backend-go/application/interfaces"
	"backend-go/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProductHanlder(
	w http.ResponseWriter,
	r *http.Request,
	repo repository.ProductRepository) error {

	var requestDto CreateProductRequestDto

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&requestDto)

	if err != nil {
		return err
	}

	if err = requestDto.Validate(); err != nil {
		return err
	}

	res, err := repo.AddProduct(
		&domain.Product{
			Model:   requestDto.Model,
			Company: requestDto.Company,
			Price:   requestDto.Price,
		})

	if err != nil {
		return err
	}

	// TODO Сделать универсальную обработку ответов на запросы ( Envelope )
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Продукт успешно создан id: %d", *res)

	return nil
}
