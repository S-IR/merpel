package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgraph-io/badger"
	"github.com/google/uuid"
	"github.com/s-ir/merpel/db"
)

func Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing API key", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		parsedId, err := uuid.Parse(parts[1])
		if err != nil {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}
		// Validate the API key
		user, err := db.GetUser(parsedId)
		if err == badger.ErrKeyNotFound {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
