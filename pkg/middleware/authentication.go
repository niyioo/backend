package middleware

import (
	"net/http"
	"strings"

	"github.com/niyioo/backend/pkg/handlers"
)

// AuthenticateMiddleware authenticates users using JWT
func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		tokenPart := splitted[1]
		user, err := handlers.ValidateToken(tokenPart)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		// Attach the user to the request context for use in the handler
		r = r.WithContext(handlers.SetUserContext(r.Context(), user))

		next.ServeHTTP(w, r)
	})
}
