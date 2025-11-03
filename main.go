// @title Backend-Go API Title
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /api/v1
// @schemes http

package main

import (
	"fmt"
	"net/http"

	_ "backend-go/docs"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	httpSwagger "github.com/swaggo/http-swagger"
)

// init вызывается до main()
func init() {
	// Загружаем значения из .env в систему
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func handleSwagger() {
	// Swagger UI будет доступен по http://localhost:8080/swagger/index.html
	http.Handle("/swagger/", httpSwagger.WrapHandler)
}

func main() {
	handleSwagger()
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/products", CreateProductHanlder).Methods("POST")

	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
