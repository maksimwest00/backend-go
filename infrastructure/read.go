package infrastructure

import (
	"backend-go/domain"
	"database/sql"
	"fmt"
	"os"
)

func Read() {
	connStr := os.Getenv("DB_CONNECT")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(
		`SELECT * from Products`)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []domain.Product{}

	for rows.Next() {
		p := domain.Product{}
		err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)

		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	for _, p := range products {
		fmt.Println(p.Id, p.Model, p.Company, p.Price)
	}
}
