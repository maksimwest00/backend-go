package main

import (
	api "backend-go/api"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// init вызывается до main()
func init() {
	// Загружаем значения из .env в систему
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func FillEmptyBodyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, пустое ли тело
		bodyBytes, _ := io.ReadAll(r.Body)

		if len(bodyBytes) == 0 {
			// Создаем тестовый объект
			testProduct := api.CreateProductDto{
				Id:      1,
				Model:   "Test Model",
				Company: "Test Company",
				Price:   1000,
			}

			// Конвертируем в JSON
			jsonData, err := json.Marshal(testProduct)
			if err != nil {
				http.Error(w, "Error creating test data", http.StatusInternalServerError)
				return
			}

			// Заполняем тело тестовыми данными
			r.Body = io.NopCloser(bytes.NewBuffer(jsonData))
			r.ContentLength = int64(len(jsonData))

		} else {
			// Если тело не пустое - восстанавливаем как есть
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Читаем тело запроса
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading body: %v", err)
			http.Error(w, "Can't read body", http.StatusBadRequest)
			return
		}

		// Логируем
		fmt.Printf("Request: %s %s\nBody: %s", r.Method, r.URL.Path, string(body))

		// Восстанавливаем тело для следующих обработчиков
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()

	// router.Use(FillEmptyBodyMiddleware)
	// router.Use(LoggingMiddleware)

	// router.HandleFunc("/about/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	id := vars["id"]
	// 	response := fmt.Sprintf("Product %s", id)
	// 	fmt.Fprint(w, response)
	// })
	// router.HandleFunc("/create", app.CreateProductHanlder).Methods("POST")

	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
