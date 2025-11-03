package repository

import "backend-go/domain"

type ProductRepository interface {
	AddProduct(*domain.Product) (*int, error)
	GetProductById(*int) (*domain.Product, error)
}
