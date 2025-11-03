package persistence

import (
	repository "backend-go/application/interfaces"
	"database/sql"
	"os"
)

type Repositories struct {
	Product repository.ProductRepository
	db      *sql.DB
}

func NewRepositories() (*Repositories, error) {
	connStr := os.Getenv("DB_CONNECT")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return &Repositories{
		Product: NewProductRepository(db),
		db:      db,
	}, nil
}

func (s *Repositories) Close() error {
	return s.db.Close()
}
