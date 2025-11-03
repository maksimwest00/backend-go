package infrastructure

import (
	"database/sql"
	"fmt"
	"os"
)

func Create(
	model string,
	company string,
	price int) {
	connStr := os.Getenv("DB_CONNECT")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec(
		`
		INSERT INTO Products (model, company, price)
		VALUES ('iPhone 17', $1, $2)
		`,
		"Apple", 72000)

	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println(count)
}
