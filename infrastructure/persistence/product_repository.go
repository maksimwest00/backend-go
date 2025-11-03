package persistence

import (
	repository "backend-go/application/interfaces"
	"backend-go/domain"
	"database/sql"
)

type ProductRepo struct {
	db *sql.DB
}

var _ repository.ProductRepository = &ProductRepo{}

// AddProduct implements repository.ProductRepository.
func (r *ProductRepo) AddProduct(p *domain.Product) (*int, error) {

	var id *int

	err := r.db.QueryRow(
		`INSERT INTO Products (model, company, price)
		VALUES ($1, $2, $3)
		RETURNING id`,
		p.Model, p.Company, p.Price,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return id, err
}

// GetProductById implements repository.ProductRepository.
func (p *ProductRepo) GetProductById(*int) (*domain.Product, error) {
	panic("unimplemented")
}

func NewProductRepository(db *sql.DB) *ProductRepo {
	return &ProductRepo{db}
}
