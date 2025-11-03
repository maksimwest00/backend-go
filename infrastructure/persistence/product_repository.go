package persistence

import (
	repository "backend-go/application/interfaces"
	"backend-go/domain"
	"database/sql"
	"fmt"
)

type ProductRepo struct {
	db *sql.DB
}

var _ repository.ProductRepository = &ProductRepo{}

func NewProductRepository(db *sql.DB) *ProductRepo {
	return &ProductRepo{db}
}

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
func (r *ProductRepo) GetProductById(id *int) (*domain.Product, error) {
	row := r.db.QueryRow(
		`SELECT id, model, company, price FROM Products
		WHERE id = $1`, id)

	p := &domain.Product{}
	err := row.Scan(&p.Id, &p.Model, &p.Company, &p.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product with id %d not found", id)
		}
		return nil, err
	}

	return p, nil
}

func (r *ProductRepo) GetProducts() (*[]domain.Product, error) {
	rows, err := r.db.Query(
		`SELECT * from Products`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []domain.Product{}

	for rows.Next() {
		p := domain.Product{}
		err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)

		if err != nil {
			continue
		}

		products = append(products, p)
	}

	return &products, nil
}
