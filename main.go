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

	"backend-go/application"
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

// @Tags Продукты
// @Summary Создать продукт
// @Description Создать продукт
// @Accept json
// @Produce json
// @Param JSON body api.CreateProductRequestDto true "CreateProductRequestDto"
// @Success 200 {object} string
// @Router /products [POST]
func CreateProductHanlder(
	w http.ResponseWriter,
	r *http.Request) {

	err := application.CreateProductHanlder(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	handleSwagger()
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/products", CreateProductHanlder).Methods("POST")

	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
