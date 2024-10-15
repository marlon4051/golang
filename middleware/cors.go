package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

// CORSMiddleware for cors
func CORSMiddleware(handler http.Handler) http.Handler {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"}, // Cambia por el dominio del frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	return corsHandler.Handler(handler)
}
