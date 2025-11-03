package application

import (
	"backend-go/api"
	"backend-go/infrastructure"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProductHanlder(
	w http.ResponseWriter,
	r *http.Request) error {

	var requestDto api.CreateProductRequestDto

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&requestDto)

	if err != nil {
		return err
	}

	if err = requestDto.Validate(); err != nil {
		return err
	}

	res, err := infrastructure.Create(requestDto)

	if err != nil {
		return err
	}

	// TODO Сделать универсальную обработку ответов на запросы ( Envelope )
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Продукт успешно создан id: %d", *res)

	return nil
}
