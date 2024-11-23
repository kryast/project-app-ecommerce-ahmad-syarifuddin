package middleware

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

func CheckLoginMiddleware(db *sql.DB, logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				logger.Warn("Authorization header missing")
				http.Error(w, "Authorization header missing", http.StatusUnauthorized)
				return
			}

			token := extractTokenFromHeader(authHeader)
			if token == "" {
				logger.Warn("Invalid Authorization header format")
				http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
				return
			}

			valid, err := isValidToken(db, token)
			if err != nil {
				logger.Error("Error checking token in database", zap.Error(err))
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			if !valid {
				logger.Warn("Invalid token")
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func extractTokenFromHeader(authHeader string) string {
	parts := strings.Split(authHeader, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}
	return ""
}

func isValidToken(db *sql.DB, token string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM auths WHERE token = $1`
	err := db.QueryRow(query, token).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error querying database: %v", err)
	}

	return count > 0, nil
}
