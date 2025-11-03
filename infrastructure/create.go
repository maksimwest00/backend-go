package infrastructure

import (
	"backend-go/api"
	"database/sql"
	"os"
)

func Create(requestDto api.CreateProductRequestDto) (*int, error) {
	connStr := os.Getenv("DB_CONNECT")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	var id *int

	err = db.QueryRow(
		`INSERT INTO Products (model, company, price) VALUES ($1, $2, $3) RETURNING id`,
		requestDto.Model, requestDto.Company, requestDto.Price,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return id, err
}
