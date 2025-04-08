package middleware

import(
	"os"
	"net/http"
	"strings"
	"context"

	"github.com/golang-jwt/jwt/v5"
)


func AuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHead := r.Header.Get("Authorization")
		if authHead == "" || !strings.HasPrefix(authHead, "Bearer "){
			http.Error(w, "Unauthorised", http.StatusUnauthorized)
		}

		tokenStr := strings.TrimPrefix(authHead, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err!=nil || !token.Valid{
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
		}

		ctx := context.WithValue(r.Context(), "user", token.Claims)
		next.ServeHTTP(w, r.WithContext(ctx))

	})	
}