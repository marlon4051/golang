package middleware

// import (
// 	"context"
// 	"net/http"
// 	"os"
// 	"strings"

// 	"github.com/dgrijalva/jwt-go"
// )

// var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// type Claims struct {
// 	UserID int `json:"user_id"`
// 	jwt.StandardClaims
// }

// func JwtMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		authHeader := r.Header.Get("Authorization")
// 		if authHeader == "" {
// 			http.Error(w, "Need token", http.StatusUnauthorized)
// 			return
// 		}

// 		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
// 		claims := &Claims{}
// 		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 			return jwtKey, nil
// 		})

// 		if err != nil || !token.Valid {
// 			http.Error(w, "Invalid token", http.StatusUnauthorized)
// 			return
// 		}

// 		// add user id to the claims context
// 		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
