package infrastructure

import (
	"database/sql"
	"fmt"
	"os"
)

func Update(
	id int,
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
		`UPDATE Products 
		 SET model = $1,
		 company = $2,
		 price = $3 
		 WHERE id = $4;`,
		model, company, price, id)

	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println(count)
}
