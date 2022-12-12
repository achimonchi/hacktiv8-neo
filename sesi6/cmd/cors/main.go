package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	index := http.HandlerFunc(Index)
	http.Handle("/index", CORS(index))

	http.ListenAndServe(":5555", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	WriteJson(w, map[string]interface{}{
		"message": "success",
	})
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		corsData := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "OPTIONS", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"token"},
			Debug:          true,
		})

		next = corsData.Handler(next)
		next.ServeHTTP(w, r)
	})
}

func WriteJson(w http.ResponseWriter, data interface{}) {

	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
	// w.Header().Set("Access-Control-Allow-Headers", "token")
	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
