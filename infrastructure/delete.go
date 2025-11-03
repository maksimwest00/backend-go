package infrastructure

import (
	"database/sql"
	"fmt"
	"os"
)

func Delete(id int) {
	connStr := os.Getenv("DB_CONNECT")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec(
		`DELETE FROM Products
		 WHERE id = $1`,
		id)

	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println(count)
}
