package createProduct

import (
	"errors"
)

type CreateProductRequestDto struct {
	Model   string `json:"model"`
	Company string `json:"company"`
	Price   int    `json:"price"`
}

func (requestDto CreateProductRequestDto) Validate() error {
	if len(requestDto.Model) == 0 {
		return errors.New("поле model не заполнено")
	}
	if len(requestDto.Company) == 0 {
		return errors.New("поле company не заполнено")
	}
	if requestDto.Price < 0 {
		return errors.New("цена ниже нуля")
	}
	return nil
}
