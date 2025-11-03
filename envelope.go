package main

// DeepSeek вариант

import (
	"encoding/json"
	"net/http"
	"time"
)

// Response стандартная структура ответа
type Response struct {
	Status    string      `json:"status"`          // "success" или "error"
	Data      interface{} `json:"data,omitempty"`  // полезные данные
	Error     string      `json:"error,omitempty"` // сообщение об ошибке
	Timestamp time.Time   `json:"timestamp"`       // время ответа
	Path      string      `json:"path,omitempty"`  // путь запроса
}

// Success создает успешный ответ
func Success(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	response := Response{
		Status:    "success",
		Data:      data,
		Timestamp: time.Now(),
		Path:      r.URL.Path,
	}

	writeJSON(w, statusCode, response)
}

// Error создает ответ с ошибкой
func Error(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	response := Response{
		Status:    "error",
		Error:     message,
		Timestamp: time.Now(),
		Path:      r.URL.Path,
	}

	writeJSON(w, statusCode, response)
}

func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
