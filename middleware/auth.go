package middleware

import (
	"context"
	"net/http"
	"strings"

	"chai/database/sqlc"
	"chai/repos"

	"github.com/rs/zerolog/log"
)

func MustGetUserFromContext(ctx context.Context) *sqlc.GetUserFromSessionContextRow {
	user, ok := ctx.Value(userCtxKey).(*sqlc.GetUserFromSessionContextRow)
	if !ok {
		panic("user not present in context")
	}
	return user
}

func APIAuthorization(repo repos.SessionRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			log.Info().Str("auth header", authHeader).Send()

			if authHeader == "" {
				http.Error(w, "authorization header missing", http.StatusForbidden)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid authorization header format", http.StatusBadRequest)
				return
			}
			token := parts[1]

			user, err := repo.GetUserFromSessionContext(token)
			if err != nil {
				log.Warn().Err(err).Send()
				http.Error(w, "unable to retrieve session", http.StatusInternalServerError)
				return
			}

			if user == nil {
				http.Error(w, "access denied", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
